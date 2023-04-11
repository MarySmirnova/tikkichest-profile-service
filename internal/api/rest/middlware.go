package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/xerrors"
	"github.com/uptrace/bunrouter"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func (s *Server) writeResponseError(w http.ResponseWriter, r *http.Request, err error, code int) {
	/*	log.Debugf("api response error: path: %v, request_time: %v, code: %v, request_id: %v",
		r.URL.Path,
		time.Now().Format("2006-01-02 15:04:05.000000"),
		code,
		r.Context().Value(ContextReqIDKey))
	*/
	s.writeCodeHeader(w, code)
	_ = json.NewEncoder(w).Encode(err.Error())
}

func (s *Server) writeCodeHeader(w http.ResponseWriter, code int) {
	w.Header().Add("Code", strconv.Itoa(code))
	w.WriteHeader(code)
}

func (s *Server) setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *Server) errorProcessingWrap(handler func(http.ResponseWriter, *http.Request) (interface{}, error)) bunrouter.HandlerFunc {
	return bunrouter.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 10<<10)
				n := runtime.Stack(buf, false)
				fmt.Fprintf(os.Stderr, "panic: %v\n\n%s", err, buf[:n])

				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			}
		}()

		resp, err := handler(w, r)
		var code int

		if _, ok := err.(*xerrors.ErrHTTP); !ok {
			switch {
			case err == nil:
			case errors.Is(err, xerrors.ErrForbidden):
				err = xerrors.NewErrHTTP(err, http.StatusForbidden)
			case errors.Is(err, xerrors.ErrInvalidReqBody):
				err = xerrors.NewErrHTTP(err, http.StatusBadRequest)
			default:
				err = xerrors.NewErrHTTP(err, http.StatusInternalServerError)
			}
		}

		switch e := err.(type) {
		case nil:
			if resp == nil {
				code = http.StatusNoContent
			} else {
				code = http.StatusOK
			}
		case *xerrors.ErrHTTP:
			code = e.Code
			resp = e.Error()
		default:
			code = http.StatusInternalServerError
			resp = e.Error()
		}

		s.setHeaders(w)

		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(resp)
	})
}

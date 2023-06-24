package rest

import (
	"context"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/auth"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/config"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/model"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

type Storage interface {
	GetProfile(username string) (*model.Profile, error)
}

type Server struct {
	db         Storage
	auth       *auth.Auth
	httpServer *http.Server
	notify     chan error
}

func NewServer(cfg config.Server, db Storage, auth *auth.Auth) *Server {
	s := &Server{
		db:     db,
		auth:   auth,
		notify: make(chan error, 1),
	}

	handler := bunrouter.New()

	//TODO: endpoints

	swagHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"))
	handler.GET("/swagger/:*", bunrouter.HTTPHandlerFunc(swagHandler))

	s.httpServer = &http.Server{
		Addr:         cfg.Listen,
		Handler:      handler,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	return s
}

func (s *Server) Start() {
	log.Printf("start server on port %s", s.httpServer.Addr)

	go func() {
		s.notify <- s.httpServer.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Shutdown(ctx context.Context) {
	_ = s.httpServer.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.httpServer
}

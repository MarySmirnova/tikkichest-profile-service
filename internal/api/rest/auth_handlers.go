package rest

import (
	"encoding/json"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/request"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/response"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/auth"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/xerrors"
	"net/http"
	"strconv"
)

// LoginHandler
// @Summary Login
// @Tags Authorize
// @Description User login
// @Accept  json
// @Produce  json
// @Param input body request.Login true "login"
// @Success 200 {object} response.Token
// @Router /login [post]
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	var login request.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		return nil, xerrors.ErrInvalidReqBody
	}

	if err := login.Validate(); err != nil {
		return nil, xerrors.NewErrHTTP(err, http.StatusBadRequest)
	}

	profile, err := s.db.GetProfile(login.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile from DB: %w", err)
	}

	token, err := s.auth.GenerateToken(profile)
	if err != nil {
		return nil, err
	}

	refresh, err := s.auth.GenerateRefresh(r.Context(), profile)
	if err != nil {
		return nil, err
	}

	return response.Token{
		Username:     token.Username,
		AccessToken:  token.AccessToken,
		ExpiresAt:    token.ExpTime.Unix(),
		RefreshToken: refresh.RefToken,
	}, nil
}

// RefreshHandler
// @Summary Refresh
// @Tags Authorize
// @Description Refresh token
// @Accept  json
// @Produce  json
// @Param input body request.Refresh true "refresh"
// @Success 200 {object} response.Token
// @Router /refresh [post]
func (s *Server) RefreshHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	var ref request.Refresh
	if err := json.NewDecoder(r.Body).Decode(&ref); err != nil {
		return nil, xerrors.ErrInvalidReqBody
	}

	if err := ref.Validate(); err != nil {
		return nil, xerrors.NewErrHTTP(err, http.StatusBadRequest)
	}

	profile, err := s.db.GetProfile(ref.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile from DB: %w", err)
	}

	refIsValid, err := s.auth.CheckRefresh(r.Context(), &auth.Refresh{
		Username: ref.Username,
		RefToken: ref.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	if !refIsValid {
		return nil, xerrors.ErrForbidden
	}

	token, err := s.auth.GenerateToken(profile)
	if err != nil {
		return nil, err
	}

	refresh, err := s.auth.GenerateRefresh(r.Context(), profile)
	if err != nil {
		return nil, err
	}

	return response.Token{
		Username:     token.Username,
		AccessToken:  token.AccessToken,
		ExpiresAt:    token.ExpTime.Unix(),
		RefreshToken: refresh.RefToken,
	}, nil
}

// AuthorizeHandler
// @Summary Authorize request
// @Tags Authorize
// @Description Token must be in the Authorize header.
// @Description If the token is invalid, return 403.
// @Description If the token is ok, return profile info for claims in the custom headers: X-Username and X-Is-Creator.
// @Accept  json
// @Produce  json
// @Success 204
// @Router /auth [head]
func (s *Server) AuthorizeHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	token := bearerAuthHeader(r.Header.Get(authHeader))

	claims, err := s.auth.ParseToken(token)
	if err != nil {
		return nil, err
	}

	w.Header().Set(usernameHeader, claims.Username)
	w.Header().Set(creatorHeader, strconv.FormatBool(claims.IsCreator))

	return nil, nil
}

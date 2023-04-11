package rest

import (
	"encoding/json"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/request"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/response"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/auth"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/xerrors"
	"net/http"
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

func (s *Server) AuthorizeHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return nil, nil
}

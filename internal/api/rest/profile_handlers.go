package rest

import (
	"encoding/json"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/request"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/api/rest/response"
	. "github.com/MarySmirnova/tikkichest-profile-service/internal/errors"
	"github.com/uptrace/bunrouter"
	"net/http"
)

// GetProfileHandler
// @Summary Get profile
// @Tags Profile
// @Description Gets user profile.
// @Accept  json
// @Produce  json
// @Param username path string true "username"
// @Success 200 {object} response.Profile
// @Router /profile/{username} [get]
func (s *Server) GetProfileHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	params := bunrouter.ParamsFromContext(r.Context())
	username, ok := params.Get("username")
	if !ok {
		return nil, fmt.Errorf("wrong username: %w", ErrInvalidInputData)
	}

	profile, err := s.db.GetProfile(r.Context(), username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile from DB: %w", err)
	}

	return response.ProfileFromDBModel(profile), nil
}

// GetProfilePageHandler
// @Summary Get profile page
// @Tags Profile
// @Description Gets the user profile page.
// @Description Takes the page marker in the "from" query parameter, if it is not set, then returns the first page.
// @Description Takes the number of elements on the page in the "limit" query parameter, if it is not set, then default value will be set.
// @Accept  json
// @Produce  json
// @Param from query string false "from"
// @Param limit query string false "limit"
// @Success 200 {object} response.ProfilePage
// @Router /profile [get]
func (s *Server) GetProfilePageHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	p, err := s.pageInfo(r)
	if err != nil {
		return nil, fmt.Errorf("wrong page info: %w", err)
	}

	profiles, nextFrom, err := s.db.PageOfProfiles(r.Context(), uint(p.limit), p.from)
	if err != nil {
		return nil, fmt.Errorf("failed to get page of user profiles from DB: %w", err)
	}

	return response.ProfilePageFromDBModel(profiles, p.limit, nextFrom), nil
}

// CreateProfileHandler
// @Summary Add new user profile.
// @Tags Profile
// @Description Creates a new user profile.
// @Accept  json
// @Produce  json
// @Param input body request.Profile true "profile"
// @Success 204
// @Router /profile [post]
func (s *Server) CreateProfileHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var profileReq request.Profile
	if err := json.NewDecoder(r.Body).Decode(&profileReq); err != nil {
		return nil, ErrInvalidReqBody
	}
	defer r.Body.Close()

	if err := profileReq.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request data: %w", err)
	}

	dbProfile, err := profileReq.ToDBModel()
	if err != nil {
		return nil, fmt.Errorf("invalid request data: %w", err)
	}

	if err := s.db.CreateProfile(r.Context(), dbProfile); err != nil {
		return nil, fmt.Errorf("failed to create user profile in DB: %w", err)
	}

	s.writeCodeHeader(w, http.StatusNoContent)
	return nil, nil
}

// UpdateProfileHandler
// @Summary Update user profile.
// @Tags Profile
// @Description Updates user profile.
// @Accept  json
// @Produce  json
// @Param username path string true "username"
// @Param input body request.ProfileUpdate true "profile"
// @Success 204
// @Router /profile/{username} [patch]
func (s *Server) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	params := bunrouter.ParamsFromContext(r.Context())
	username, ok := params.Get("username")
	if !ok {
		return nil, fmt.Errorf("wrong username: %w", ErrInvalidInputData)
	}

	var profileReq request.ProfileUpdate
	if err := json.NewDecoder(r.Body).Decode(&profileReq); err != nil {
		return nil, ErrInvalidReqBody
	}

	if err := profileReq.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request data: %w", err)
	}

	dbProfile, err := profileReq.ToDBModel()
	if err != nil {
		return nil, fmt.Errorf("invalid request data: %w", err)
	}

	if err := s.db.UpdateProfile(r.Context(), username, dbProfile); err != nil {
		return nil, fmt.Errorf("failed to update user profile in DB: %w", err)
	}

	return nil, nil
}

// DeleteProfileHandler
// @Summary Delete user profile.
// @Tags Profile
// @Description Deletes a user profile.
// @Accept  json
// @Produce  json
// @Param username path string true "username"
// @Success 204
// @Router /profile/{username} [delete]
func (s *Server) DeleteProfileHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	params := bunrouter.ParamsFromContext(r.Context())
	username, ok := params.Get("username")
	if !ok {
		return nil, fmt.Errorf("wrong username: %w", ErrInvalidInputData)
	}

	if err := s.db.DeleteProfile(r.Context(), username); err != nil {
		return nil, fmt.Errorf("failed to delete user profile in DB: %w", err)
	}

	return nil, nil
}

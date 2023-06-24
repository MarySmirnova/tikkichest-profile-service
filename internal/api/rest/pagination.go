package rest

import (
	"fmt"
	. "github.com/MarySmirnova/tikkichest-profile-service/internal/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

const DefaultProfilesPerPage = 20

type pageParameters struct {
	from  string
	limit int
}

func (s *Server) pageInfo(r *http.Request) (*pageParameters, error) {
	limit := DefaultProfilesPerPage

	limitString := r.FormValue("limit")

	if limitString != "" {
		l, err := strconv.Atoi(limitString)
		if err != nil {
			return nil, fmt.Errorf("'limit' must be an integer greater than zero: %w", ErrInvalidParameterFormat)
		}
		limit = l
	}

	if limit <= 0 {
		return nil, fmt.Errorf("'limit' must be an integer greater than zero: %w", ErrInvalidParameterFormat)
	}

	from := r.FormValue("from")

	if from == "" {
		from = primitive.NilObjectID.Hex()
	}

	_, err := primitive.ObjectIDFromHex(from)
	if err != nil {
		return nil, fmt.Errorf("invalid marker 'from': %w", ErrInvalidParameterFormat)
	}

	return &pageParameters{
		from:  from,
		limit: limit,
	}, nil
}

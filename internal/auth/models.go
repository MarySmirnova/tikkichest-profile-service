package auth

import (
	"github.com/MarySmirnova/tikkichest-profile-service/internal/db/model"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username  string `json:"username"`
	IsCreator bool   `json:"is_creator"`
	jwt.StandardClaims
}

type Refresh struct {
	Username string
	RefToken string
}

func (r *Refresh) fillFromDB(dbRefresh *model.Refresh) {
	r.Username = dbRefresh.Username
	r.RefToken = dbRefresh.RefToken
}

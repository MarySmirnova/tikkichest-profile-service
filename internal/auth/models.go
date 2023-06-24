package auth

import (
	"github.com/MarySmirnova/tikkichest-profile-service/internal/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username  string `json:"username"`
	IsCreator bool   `json:"is_creator"`
	jwt.StandardClaims
}

type Token struct {
	Username    string
	ExpTime     time.Time
	AccessToken string
}

type Refresh struct {
	Username string
	RefToken string
}

func (r *Refresh) fillFromDB(dbRefresh *model.Refresh) {
	r.Username = dbRefresh.Username
	r.RefToken = dbRefresh.RefToken
}

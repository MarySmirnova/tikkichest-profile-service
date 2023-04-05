package auth

import (
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/db/model"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (a *Auth) GenerateToken(profile *model.Profile) (string, error) {
	timeNow := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodRS384,
		&Claims{
			Username:  profile.Username,
			IsCreator: profile.IsCreator,
			StandardClaims: jwt.StandardClaims{
				Issuer:    a.issuer,
				IssuedAt:  timeNow.Unix(),
				ExpiresAt: timeNow.Add(a.accessTokenTime).Unix(),
			}})

	tokenStr, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token, %w", err)
	}

	return tokenStr, nil
}

func (a *Auth) ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.publicKey, nil
	})
	if err != nil {
		return nil, errors.ErrForbidden
	}
	return claims, nil
}

func (a *Auth) GenerateRefresh() {
	//сгенерировать рефреш, положить его в редис (с временем протухания), отдать
	//если у юзера уже есть рефреш, заменить
}

func (a *Auth) CheckRefresh() bool {
	//проверить есть ли рефреш в редисе
	return true
}

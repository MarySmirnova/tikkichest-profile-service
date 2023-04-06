package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/db/model"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/xerrors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

// GenerateToken генерирует новый токен для пользователя.
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

// ParseToken проверяет валидность токена. Если токен валиден, возвращает Claims.
// Если токен не валиден, вернется ошибка xerrors.ErrForbidden.
func (a *Auth) ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.publicKey, nil
	})
	if err != nil {
		return nil, xerrors.ErrForbidden
	}
	return claims, nil
}

// GenerateRefresh генерирует новый рефреш токен для пользователя.
func (a *Auth) GenerateRefresh(ctx context.Context, profile *model.Profile) (*Refresh, error) {
	timeNow := time.Now()

	refreshDB := &model.Refresh{
		Username:  profile.Username,
		RefToken:  uuid.NewString(),
		IssuedAt:  timeNow.Unix(),
		ExpiresAt: timeNow.Add(a.refreshTokenTime).Unix(),
	}

	err := a.db.CreateRefresh(ctx, refreshDB, a.refreshTokenTime)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token to db: %w", err)
	}

	refresh := new(Refresh)
	refresh.fillFromDB(refreshDB)

	return refresh, nil
}

// CheckRefresh проверяет, валиден ли рефреш токен пользователя.
func (a *Auth) CheckRefresh(ctx context.Context, refresh *Refresh) (bool, error) {
	refreshDB, err := a.db.GetRefresh(ctx, refresh.Username)
	if err != nil {
		if errors.Is(err, xerrors.ErrMissingFromDB) {
			return false, nil
		}
		return false, fmt.Errorf("failed to get refresh token from db: %w", err)
	}

	timeNow := time.Now()
	timeExp := time.Unix(refreshDB.ExpiresAt, 0)
	if timeExp.Before(timeNow) {
		return false, nil
	}

	if refreshDB.RefToken != refresh.RefToken {
		return false, nil
	}

	return true, nil
}

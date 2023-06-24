package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/config"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/model"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type RefreshStorage interface {
	CreateRefresh(ctx context.Context, refresh *model.Refresh, expiration time.Duration) error
	GetRefresh(ctx context.Context, username string) (*model.Refresh, error)
}

type Auth struct {
	db               RefreshStorage
	privateKey       *rsa.PrivateKey
	publicKey        *rsa.PublicKey
	accessTokenTime  time.Duration
	refreshTokenTime time.Duration
	issuer           string
}

func New(cfg config.Auth, db RefreshStorage) (*Auth, error) {
	a := &Auth{
		db:               db,
		accessTokenTime:  cfg.AccessTokenTime,
		refreshTokenTime: cfg.RefreshTokenTime,
		issuer:           cfg.Issuer,
	}

	privateKeyRaw, err := os.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}
	publicKeyRaw, err := os.ReadFile(cfg.PublicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	a.privateKey = privateKey
	a.publicKey = publicKey

	return a, nil
}

package main

import (
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/auth"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/config"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/db/model"
	"time"
)

func main() {
	a, err := auth.New(config.Auth{
		Issuer:           "test",
		RefreshTokenTime: 3 * time.Minute,
		AccessTokenTime:  15 * time.Minute,
		PrivateKeyFile:   "id_rsa",
		PublicKeyFile:    "id_rsa.pub",
	}, "")
	if err != nil {
		fmt.Println(err)
	}

	token, err := a.GenerateToken(&model.Profile{
		Username: "user1", IsCreator: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("creator:\n%s\n", token)

	claims, err := a.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *claims)
}

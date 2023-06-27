package request

import (
	"fmt"
	. "github.com/MarySmirnova/tikkichest-profile-service/internal/errors"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/model"
	"github.com/MarySmirnova/tikkichest-profile-service/pkg/hashpass"
	"time"
)

type Profile struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	Town      string `json:"town"`
	Password  string `json:"password"`
	IsCreator bool   `json:"is_creator"`
}

func (p *Profile) ToDBModel() (*model.Profile, error) {
	hashPassword, err := hashpass.CreatePasswordHash(p.Password)
	if err != nil {
		return nil, fmt.Errorf("create password hash error: %w", err)
	}

	timeNow := time.Now().Unix()

	return &model.Profile{
		Username:     p.Username,
		Name:         p.Name,
		Email:        p.Email,
		Phone:        p.Phone,
		HashPassword: hashPassword,
		IsCreator:    p.IsCreator,
		CreateTime:   timeNow,
		ChangeTime:   timeNow,
		Location: model.Location{
			Country: p.Country,
			Town:    p.Town,
		},
	}, nil
}

func (p *Profile) Validate() error {
	err := NewValidError()

	if p.Username == "" {
		err.Add("username", "username must be filled")
	}

	if p.Name == "" {
		err.Add("name", "name must be filled")
	}

	if p.Password == "" {
		err.Add("password", "password must be filled")
	}

	return err.Check()
}

type ProfileUpdate struct {
	Username *string `json:"username,omitempty"`
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Country  *string `json:"country,omitempty"`
	Town     *string `json:"town,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (p *ProfileUpdate) Validate() error {
	err := NewValidError()

	//TODO: add validate parameters

	return err.Check()
}

func (p *ProfileUpdate) ToDBModel() (*model.ProfileUpdate, error) {
	timeNow := time.Now().Unix()

	profile := &model.ProfileUpdate{
		Username:   p.Username,
		Name:       p.Name,
		Email:      p.Email,
		Phone:      p.Phone,
		ChangeTime: timeNow,
	}

	if p.Password != nil {
		hashPassword, err := hashpass.CreatePasswordHash(*p.Password)
		if err != nil {
			return nil, fmt.Errorf("create password hash error: %w", err)
		}
		profile.HashPassword = &hashPassword
	}

	if p.Country != nil || p.Town != nil {
		profile.Location = &model.LocationUpdate{
			Country: p.Country,
			Town:    p.Town,
		}
	}

	return profile, nil
}

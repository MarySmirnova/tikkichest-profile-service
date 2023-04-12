package request

import "fmt"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *Login) Validate() error {
	if l.Username == "" {
		return fmt.Errorf("username must be filled")
	}

	if len([]rune(l.Password)) < 6 {
		return fmt.Errorf("too short password")
	}
	return nil
}

type Refresh struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}

func (r *Refresh) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("username must be filled")
	}

	return nil
}

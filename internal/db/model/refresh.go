package model

type Refresh struct {
	Username  string
	RefToken  string
	IssuedAt  int64
	ExpiresAt int64
}

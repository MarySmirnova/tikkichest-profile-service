package response

type Token struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Username  string `json:"username"`
	IsCreator bool   `json:"is_creator"`
}

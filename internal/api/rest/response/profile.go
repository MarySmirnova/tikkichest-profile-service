package response

type Profile struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	Town      string `json:"town"`
	IsCreator bool   `json:"is_creator"`
}

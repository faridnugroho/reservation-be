package dto

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	NoHP     string `json:"noHp"`
	Password string `json:"password,omitempty"`
}

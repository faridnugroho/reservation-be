package dto

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password,omitempty"`
}

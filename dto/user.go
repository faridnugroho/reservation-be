package dto

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password,omitempty"`
}

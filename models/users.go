package models

type Users struct {
	CustomPublicModel
	Fullname string `gorm:"type:varchar(255);column:fullname" json:"fullname"`
	Email    string `gorm:"type:varchar(255);column:email" json:"email"`
	Password string `gorm:"type:varchar(255);column:password" json:"password"`
}

func (Users) TableName() string {
	return "users"
}

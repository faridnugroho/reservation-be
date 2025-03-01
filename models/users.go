package models

type Users struct {
	CustomPublicModel
	Fullname   string `gorm:"type:varchar(255);column:fullname" json:"fullname"`
	Email      string `gorm:"type:varchar(255);column:email" json:"email"`
	No_hp      string `gorm:"type:varchar(255);column:no_hp" json:"no_hp"`
	Password   string `gorm:"type:varchar(255);column:password" json:"password"`
	IsVerified bool   `gorm:"type:bool;default:false;column:is_verified" json:"is_verified"`
}

func (Users) TableName() string {
	return "users"
}

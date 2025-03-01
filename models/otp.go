package models

import "github.com/google/uuid"

type Otp struct {
	CustomPublicModel
	UserID  uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`
	OtpCode string    `gorm:"type:varchar(255);column:otp_code" json:"otp_code"`
}

func (Otp) TableName() string {
	return "otps"
}

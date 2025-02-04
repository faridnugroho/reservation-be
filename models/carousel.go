package models

type Carousels struct {
	CustomPublicModel
	Url    string `gorm:"type:varchar(255);column:url" json:"url"`
	Status bool   `gorm:"type:bool;default:false;column:status" json:"status"`
}

func (Carousels) TableName() string {
	return "carousels"
}

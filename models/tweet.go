package models

type Tweet struct {
	Base
	Content   string `json:"content"`
	Createdby User   `json:"createdby" gorm:"ForeignKey:UserId"`
	UserId    uint   `json:"user_id"`
}

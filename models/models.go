package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint      `gorm:"primarykey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

//GenerateTimeStamps -Generate timestamps for createdat and updatedat in string format
func GenerateTimeStamps() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// BeforeCreate - sets Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	//Create new random UUIDS
	base.UUID = uuid.New()
	//TIMESTAMPS
	t := GenerateTimeStamps()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerateTimeStamps()
	return nil
}

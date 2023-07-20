package models

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"Created"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

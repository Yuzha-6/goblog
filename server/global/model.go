package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

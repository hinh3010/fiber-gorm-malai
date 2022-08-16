package emtity

import (
	"time"

	"gorm.io/gorm"
)

type UserEmtity struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Age       uint8          `json:"age"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

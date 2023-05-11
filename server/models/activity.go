package models

import "time"

type Activities struct {
	ActivityID uint      `gorm:"primaryKey" json:"activity_id"`
	Title      string    `gorm:"type:varchar(100)" json:"title"`
	Email      string    `gorm:"type:varchar(100)" json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

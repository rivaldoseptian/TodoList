package models

import "time"

type Todo struct {
	TodoID          uint      `gorm:"primaryKey" json:"todo_id"`
	ActivityGroupID uint      `gorm:"type:integer" json:"activity_group_id"`
	Title           string    `gorm:"type:varchar(100)" json:"title"`
	IsActive        bool      `gorm:"type:boolean" json:"is_active"`
	Priority        string    `gorm:"type:varchar(100)" json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

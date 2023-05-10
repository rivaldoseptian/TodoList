package models

import "time"

type Todo struct {
	TodoID          uint      `gorm:"primaryKey" json:"todo_id"`
	ActivityGroupID uint      `gorm:"type:number" json:"activity_group_id"`
	Title           string    `gorm:"type:varchar(100)" json:"title"`
	IsActive        bool      `gorm:"type:bool" json:"is_active"`
	Priority        string    `gorm:"type:string" json:"priority"`
	CratedAt        time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

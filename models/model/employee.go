package model

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Age      int    `json:"age" gorm:"not null"`
	JobTitle string `json:"job_title" gorm:"not null"`
	Company  string `json:"company"  gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

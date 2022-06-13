package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	gorm.Model
	// AttendanceID uint64 `json:"id" gorm:"primary_key"`
	EmployeeID uint
	Date       time.Time `json:"date"`
	// CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CreateAttendanceInput struct {
	EmployeeID uint
	Date       string `json:"date"`
}

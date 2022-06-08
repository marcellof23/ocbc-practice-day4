package entity

import (
	"time"
)

type Employee struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Address   string    `json:"address"`
	Job       string    `json:"job"`
	JoinDate  time.Time `json:"joinDate"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}

type UpdateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}

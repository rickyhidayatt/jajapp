package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id          int
	Uuid        uuid.UUID
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	IsSeller    bool
	IsDriver    bool
	Nik         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsVerified  bool
	Latitude    float64
	Longitude   float64
}

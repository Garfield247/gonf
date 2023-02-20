package models

import (
	"time"
)

type Services struct {
	BaseModel
	Title       string
	FullName    string
	GameID      string
	StartAt     time.Time
	EndAt       time.Time
	LifeStartAt time.Time
	LifeEndAt   time.Time
	Creater     uint64
	Updater     uint64
}

func (table *Services) TableName() string {
	return "services"
}

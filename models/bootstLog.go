package models

import "time"

type BoostLog struct {
	BaseModel
	ServiceID uint64
	StartTime time.Time
	EndTime   time.Time
	Bots      string
}

package models

import "time"

type BaseModel struct {
	ID          uint64    `gorm:"primary_key" json:"id"`
	CreatedTime time.Time `                   json:"created_time"`
	UpdatedTime time.Time `                   json:"updated_time"`
}

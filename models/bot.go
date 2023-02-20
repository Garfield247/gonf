package models

type Bot struct {
	BaseModel
	Name   string `gorm:"unique index" json:"name"`
	HostId uint64 `                    json:"host_id"`
}

func (bot *Bot) TableName() string {
	return "bot"
}

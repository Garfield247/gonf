package models

type BotHost struct {
	BaseModel
	Name   string `gorm:"unique index" json:"name"`
	Remark string `                    json:"remark"`
}

func (bot *BotHost) TableName() string {
	return "bot"
}

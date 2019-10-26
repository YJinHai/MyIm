package models

import (
	"time"
)

type ImMsgContent struct {
	Mid         int       `xorm:"not null pk autoincr INT(11)"`
	Content     string    `xorm:"not null VARCHAR(1000)"`
	SenderId    int       `xorm:"not null INT(11)"`
	RecipientId int       `xorm:"not null INT(11)"`
	MsgType     int       `xorm:"not null INT(11)"`
	CreateTime  time.Time `xorm:"not null TIMESTAMP"`
}

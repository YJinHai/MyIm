package models

import (
	"time"
)

type ImMsgSend struct {
	MsgId      int64     `xorm:"not null pk BIGINT(20)"`
	MsgFrom    int64     `xorm:"not null BIGINT(20)"`
	MsgTo      int64     `xorm:"not null BIGINT(20)"`
	MsgContent string    `xorm:"VARCHAR(255)"`
	SendTime   time.Time `xorm:"not null TIMESTAMP"`
	MsgType    int       `xorm:"not null INT(2)"`
}

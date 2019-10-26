package models

import (
	"time"
)

type ImMsgContact struct {
	OwnerUid   int       `xorm:"not null pk INT(11)"`
	OtherUid   int       `xorm:"not null pk INT(11)"`
	Mid        int       `xorm:"not null INT(11)"`
	Type       int       `xorm:"not null INT(11)"`
	CreateTime time.Time `xorm:"not null TIMESTAMP"`
}

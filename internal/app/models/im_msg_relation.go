package models

import (
	"time"
)

type ImMsgRelation struct {
	OwnerUid   int       `xorm:"not null pk index(idx_owneruid_otheruid_msgid) INT(11)"`
	OtherUid   int       `xorm:"not null index(idx_owneruid_otheruid_msgid) INT(11)"`
	Mid        int       `xorm:"not null pk index(idx_owneruid_otheruid_msgid) INT(11)"`
	Type       int       `xorm:"not null INT(11)"`
	CreateTime time.Time `xorm:"not null TIMESTAMP"`
}

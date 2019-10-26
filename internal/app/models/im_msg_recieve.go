package models

type ImMsgRecieve struct {
	Id      int64 `xorm:"pk BIGINT(20)"`
	MsgId   int64 `xorm:"not null BIGINT(20)"`
	MsgFrom int64 `xorm:"not null BIGINT(20)"`
	MsgTo   int64 `xorm:"not null BIGINT(20)"`
	Flag    int   `xorm:"not null TINYINT(1)"`
}

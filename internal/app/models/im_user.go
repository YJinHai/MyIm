package models

type ImUser struct {
	Uid      int    `xorm:"not null pk autoincr INT(11)"`
	Username string `xorm:"not null VARCHAR(500)"`
	Password string `xorm:"not null VARCHAR(500)"`
	Email    string `xorm:"VARCHAR(250)"`
	Avatar   string `xorm:"not null VARCHAR(500)"`
	Secret   string `xorm:"VARCHAR(255)"`
}

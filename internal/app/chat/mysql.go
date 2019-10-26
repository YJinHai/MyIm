package chat

import (
	"github.com/YJinHai/MyIm/internal/app/models"
	"github.com/YJinHai/MyIm/internal/pkg/snowflake"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
	"time"
)

type Dao struct {
	engine *xorm.Engine
}

func NewChatDao(engine *xorm.Engine) *Dao {
	return &Dao{
		engine: engine,
	}
}

func (d *Dao) SaveC2CSendMsg(msg *C2CSendRequest) (*C2CSendResponse, error){
	id := snowflake.GetSnowflakeId()
	data := &models.ImMsgSend{}
	data.MsgId = id
	data.MsgFrom = msg.From
	data.MsgTo = msg.To
	data.MsgContent = msg.Content
	data.MsgType = 1
	data.SendTime = time.Now()


	if _,err := d.engine.Insert(data); err != nil{
		return nil,err
	}

	r := &C2CSendResponse{
		MsgId:id,
	}

	return r,nil
}

func (d *Dao) SaveC2CPushMsg(uid int64,msg *C2CPushRequest) (*C2CPushResponse, error){
	id := snowflake.GetSnowflakeId()
	logs.Info("SaveC2CPushMsg id:",id)
	data := &models.ImMsgRecieve{}
	data.MsgId = id
	data.MsgFrom = msg.From
	data.MsgTo = uid
	data.Flag = 0

	logs.Info("SaveC2CPushMsg data:",data)



	if _,err := d.engine.Insert(data); err != nil{
		return nil,err
	}

	r := &C2CPushResponse{
		MsgId:id,
	}

	return r,nil
}

func (d *Dao) GetC2CMsg(msg *C2CPushRequest) *C2CPushResponse {
	return nil
}

package user

import (
	"errors"
	"github.com/go-xorm/xorm"

	"github.com/YJinHai/MyIm/internal/app/models"
)

type Dao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *Dao {
	return &Dao{
		engine: engine,
	}
}

func (d *Dao) GetInfo(data *InfoRequest)  (*InfoResponse,error){
	user := &models.ImUser{Uid:data.Uid}
	has,err := d.engine.Get(user)

	if !has || err != nil{
		return nil,nil
	}

	res := &InfoResponse{
		Uid:user.Uid,
		Username:user.Username,
		Email:user.Email,
	}

	return res,nil
}

func (d *Dao) Login(data *LoginRequest) (*InfoResponse,error){
	total, err := d.engine.Count(&models.ImUser{})
	if err != nil{
		return nil,errors.New("bar must not be empty")
	}

	user := &models.ImUser{
		Uid:int(total+1),
		Email:data.Email,
		Password:data.Password,
		Username:data.Email,
	}

	_,err = d.engine.Insert(user)
	if err != nil{
		return nil,err
	}

	res := &InfoResponse{
		Uid:user.Uid,
		Username:user.Username,
	}

	return res,nil
}

func (d *Dao) Register (data *RegisterRequest) (*InfoResponse,error){

	user := &models.ImUser{
		Email:data.Email,
		Password:data.Password,
	}
	has,err := d.engine.Get(user)
	if !has || err != nil{
		return nil,nil
	}

	res := &InfoResponse{
		Uid:user.Uid,
		Username:user.Username,
		Email:user.Email,
	}

	return res,nil
}

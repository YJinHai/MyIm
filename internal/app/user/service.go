package user

import (
	"github.com/YJinHai/MyIm/internal/pkg/mysql"
)

type Service interface {
	Info(data *InfoRequest) (*InfoResponse, error)
	Login(data *LoginRequest) (*InfoResponse,error)
	Register(data *RegisterRequest) (*InfoResponse,error)
}

type userService struct {
	dao *Dao
}


func NewUserService() Service {
	return &userService{
		dao: NewUserDao(mysql.GetSelfDB()),
	}
}

func (s *userService) Info(data *InfoRequest) (*InfoResponse, error){


	return s.dao.GetInfo(data)

}

func (s *userService) Login(data *LoginRequest) (*InfoResponse,error) {

	return s.dao.Login(data)
}

func (s *userService) Register(data *RegisterRequest) (*InfoResponse,error) {

	return s.dao.Register(data)
}
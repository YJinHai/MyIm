package chat

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"log"
	"regexp"
	"strconv"

	"github.com/YJinHai/MyIm/internal/app/models"
)

var (
	exceedsAny func(map[string]Limit) error
	exceeds    func(string, Limit) error
	alfaRgx    *regexp.Regexp
	mailRgx    *regexp.Regexp
)

// Limiter represents chat service limit checker
type Limiter interface {
	Exceeds(string, Limit) error
	ExceedsAny(map[string]Limit) error
}

// New creates new websocket api
func NewAPI() *API {
	return &API{
		store: NewRedis(),
	}
}

// API represents websocket api service
type API struct {
	store Store
}

// Store represents chat store interface
type Store interface {
	Save(*Chat) error
	Get(string) (*Chat, error)
	ListChannels() ([]string, error)
	GetUnreadCount(string, string) uint64
}



func (cr *CreateReq) Bind() error {
	if !alfaRgx.MatchString(cr.Name) {
		return errors.New("name must contain only alphanumeric and underscores")
	}
	return exceeds(cr.Name, ChanLimit)
}

func (api *API) CreateChannel(req *CreateReq) (string, error) {
	ch := NewChannel(req.Name, req.IsPrivate)
	if err := api.store.Save(ch); err != nil {
		// http.Error(w, fmt.Sprintf("could not create channel: %v", err), 500)
		return "", err
	}

	return ch.Secret,nil
}



//func (r *RegisterReq) Bind() error {
//	//if !alfaRgx.MatchString(r.Uid) {
//	//	return errors.New("uid must contain only alphanumeric and underscores")
//	//}
//	if !alfaRgx.MatchString(r.Secret) {
//		return errors.New("secret must contain only alphanumeric and underscores")
//	}
//	if !mailRgx.MatchString(r.Email) {
//		return errors.New("invalid email address")
//	}
//	return exceedsAny(map[string]Limit{
//		r.Uid:           UIDLimit,
//		r.DisplayName:   DisplayNameLimit,
//		r.ChannelSecret: ChanSecretLimit,
//		r.Secret:        SecretLimit,
//		r.Channel:       ChanLimit,
//	})
//}

func (api *API) Register(req *RegisterReq) (*RegisterResp, error){
	ch, err := api.store.Get(req.Channel)
	if err != nil || ch.Secret != req.ChannelSecret {

		//http.Error(w, fmt.Sprintf("invalid secret or unexisting channel: %v", err), 500)
		log.Println("invalid secret or unexciting channel: %v", err)
		return nil,err
	}

	logs.Info("ImUser:",req)
	secret, err := ch.Register(&models.ImUser{
		Uid:         req.Uid,
		Username: req.Username,
		Email:       req.Email,
		Secret:      req.Secret,
	})

	if err != nil {
		// http.Error(w, fmt.Sprintf("error registering to channel: %v", err), 500)
		log.Println("error registering to channel: %v", err)
		return nil,err
	}

	if err = api.store.Save(ch); err != nil {
		ch.Leave(req.Uid)
		//http.Error(w, fmt.Sprintf("could not update channel membership: %v", err), 500)
		log.Println("could not update channel membership: %v", err)
		return nil,err
	}

	return &RegisterResp{Secret:secret},err

}

type unreadCountResp struct {
	Count uint64 `json:"count"`
}

func (api *API) unreadCount(uid int, chanName string) (uint64, error) {
	//if err := exceedsAny(map[string]Limit{
	//	chanName: ChanLimit,
	//	uid:      UIDLimit,
	//}); err != nil {
	//	http.Error(w, err.Error(), 400)
	//	return
	//}

	uc := api.store.GetUnreadCount(strconv.Itoa(uid), chanName)
	return  uc,nil
}

func (api *API) ListMembers(chanName, secret string) ([]*models.ImUser, error) {

	//if err := exceedsAny(map[string]Limit{
	//	chanName: ChanLimit,
	//	secret:   ChanSecretLimit,
	//}); err != nil {
	//	//http.Error(w, err.Error(), 400)
	//	return nil,err
	//}

	ch, err := api.store.Get(chanName)
	if err != nil {
		return nil,err
	}

	if ch.Secret != secret {
		//http.Error(w, "invalid secret", 500)
		return nil,err
	}

	return ch.ListMembers(), nil
}

func (api *API) ListChannels() ([]string,error){
	chans, err := api.store.ListChannels()
	if err != nil {
		return nil,err
	}

	return chans,nil
}

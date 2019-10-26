package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YJinHai/MyIm/internal/pkg/mysql"
	"net/http"
	"regexp"

	"github.com/gorilla/websocket"
	"github.com/ribice/goch"

	"github.com/YJinHai/MyIm/internal/app/broker"
	"github.com/YJinHai/MyIm/internal/app/chat"
	"github.com/YJinHai/MyIm/internal/app/ingest"
	"github.com/YJinHai/MyIm/internal/pkg/nats"
)

var (
	alfaRgx *regexp.Regexp
)

// API represents websocket api service
type API struct {
	broker   *broker.Broker
	store    ChatStore
	dbstore	 DBStore
	cache    ChatStore
	upgrader websocket.Upgrader
	rlim     Limiter
}

// Limiter represents chat service limit checker
type Limiter interface {
	ExceedsAny(map[string]goch.Limit) error
}

// NewAPI creates new websocket api
func NewAPI() *API {
	mq := nats.GetSelfMQ()
	store := chat.NewRedis()
	dbstore := chat.NewChatDao(mysql.GetSelfDB())
	api := API{
		broker: broker.New(mq, store, ingest.New(mq, store)),
		store:  store,
		dbstore: dbstore,
		cache: store,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}

	return &api
}



func (api *API) Connect(w http.ResponseWriter, r *http.Request) error {
	// 升级将HTTP服务器连接升级到WebSocket协议。
	conn, err := api.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while upgrading to ws connection: %v", err), 500)
		return err
	}

	// 向聊天室注册
	req, err := api.waitConnInit(conn)
	if err != nil {
		if err == errConnClosed {
			return err
		}
		writeErr(conn, err.Error())
		return err
	}

	// 开始接受信息
	agent := NewAgent(api.broker, api.store, api.dbstore)
	agent.HandleConn(conn, req)

	return nil
}



//func (api *API) bindReq(r *InitConReq) error {
//	if !alfaRgx.MatchString(r.Secret) {
//		return errors.New("secret must contain only alphanumeric and underscores")
//	}
//	if !alfaRgx.MatchString(r.Channel) {
//		return errors.New("channel must contain only alphanumeric and underscores")
//	}
//
//	return api.rlim.ExceedsAny(map[string]goch.Limit{
//		r.UID:     goch.UIDLimit,
//		r.Secret:  goch.SecretLimit,
//		r.Channel: goch.ChanLimit,
//	})
//}

var errConnClosed = errors.New("connection closed")

func (api *API) waitConnInit(conn *websocket.Conn) (*InitConReq, error) {
	t, wsr, err := conn.NextReader()
	if err != nil || t == websocket.CloseMessage {
		return nil, errConnClosed
	}

	var req *InitConReq

	err = json.NewDecoder(wsr).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

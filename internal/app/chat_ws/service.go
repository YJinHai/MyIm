package chat_ws

import (
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

type WS interface {

}

type ws struct {

}


//func (api *ws) Connect(w http.ResponseWriter, r *http.Request) error {
//	// 升级将HTTP服务器连接升级到WebSocket协议。
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		logs.Error(err)
//		return err
//	}
//
//
//	return nil
//}

// serveWs handles websocket requests from the peer.
func Connect(hub *Hub, w http.ResponseWriter, r *http.Request, userID string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error(err)
		return
	}
	client := NewClient(conn)


	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.readPump()
	go client.writePump()
}
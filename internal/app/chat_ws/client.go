package chat_ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)


const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

//// SendMessage is
//var (
//	addr        = flag.String("addr", "modbus:8080", "http service address")
//	SendMessage = make(chan *MessageSendByWebsocket)
//	count       = 0
//)




var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)


// Client is a middleman between the websocket connection and the hub.
type Client struct {

	// The websocket connection.
	conn *websocket.Conn

	send	chan []byte     // 待发送的数据
}

// 初始化
func NewClient(socket *websocket.Conn) *Client {

	return  &Client{
		conn:        socket,
		send:          make(chan []byte, 100),
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		//_, message, err := c.conn.ReadMessage()
		var msg MyTestMSG
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		fmt.Println(msg)
		//fmt.Println(message)

		//message :=  []byte("I am reading!")
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- &msg
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("writePump:", msg.User)

			//w, err := c.conn.NextWriter(websocket.TextMessage)
			//if err != nil {
			//	return
			//}
			reData := & MyTestResponse{
				Data: msg.User,
			}

			err := c.conn.WriteJSON(reData)
			if err != nil {
				fmt.Println("error writePump")
				return
			}

			//bytes := []byte("I am byte array !")

			//w.Write(bytes)


			// Add queued chat messages to the current websocket message.
			//n := len(c.send)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	w.Write(<-c.send)
			//}
			//
			//if err := w.Close(); err != nil {
			//	return
			//}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}





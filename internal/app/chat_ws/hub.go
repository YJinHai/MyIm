package chat_ws

import (
	"github.com/YJinHai/MyIm/internal/app/broker"
	"github.com/YJinHai/MyIm/internal/app/chat"
	"github.com/YJinHai/MyIm/internal/app/ingest"
	"github.com/YJinHai/MyIm/internal/pkg/nats"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	//// Registered clients.
	//clients map[*Client]bool
	//
	//// Inbound messages from the clients.
	//broadcast chan *MyTestMSG
	//
	//// Register requests from the clients.
	//register chan map[string]*Client
	//
	//// Unregister requests from clients.
	//unregister chan *Client
	//
	////user id
	//usersID map[string]*Client
	//
	//
	//userBroadcast chan *postUser
	//
	//userRegister chan *Server
	//
	//servers map[*Server]bool
	broker   *broker.Broker
	store    ChatStore

	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据

}

func newHub() *Hub {
	mq := nats.GetSelfMQ()
	store := NewRedis()
	return &Hub{
		//broadcast:  make(chan *MyTestMSG),
		//register:   make(chan map[string]*Client),
		//unregister: make(chan *Client),
		//clients:    make(map[*Client]bool),
		//usersID:    make(map[string]*Client),
		//
		//userBroadcast:	make(chan *postUser),
		//userRegister:	make(chan *Server),
		//servers:	make(map[*Server]bool),
		broker: broker.New(mq, store, ingest.New(mq, store)),
		store:  store,

	}
}

func (h *Hub) run() {
	var clientUser map[string]*Client
	for {
		select {
		case serverUser := <-h.userRegister:
			h.servers[serverUser] = true

		case message := <-h.userBroadcast:
			for server := range h.servers {
				select {
				case server.send <- message:
				default:
					close(server.send)
					delete(h.servers, server)
				}
			}

		case clientUser = <-h.register:
			for k,v := range clientUser {
				h.clients[v] = true
				h.usersID[k] = v
				fmt.Println("client:",k,"-",v)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				fmt.Println("删除前：",h.usersID)
				for k,v := range h.usersID{
					if v == client{
						delete(h.usersID, k)

					}
				}
				fmt.Println("删除后：",h.usersID)
				close(client.send)

			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

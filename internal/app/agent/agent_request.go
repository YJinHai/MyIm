package agent

import (
	"github.com/vmihailenco/msgpack"
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

type InitConReq struct {
	Channel string  `json:"channel"`
	UID     int  `json:"uid"`
	Secret  string  `json:"secret"` // User secret
	LastSeq *uint64 `json:"last_seq"`
}

// Message represents chat message
type Message struct {
	Meta     map[string]string `json:"meta"`
	Time     int64             `json:"time"`
	Seq      uint64            `json:"seq"`
	Text     string            `json:"text"`
	FromUID  int            `json:"from_uid"`
	FromName string            `json:"from_name"`
}

// DecodeMsg tries to decode binary formatted message in b to Message
func DecodeMsg(b []byte) (*Message, error) {
	var msg Message
	err := msgpack.Unmarshal(b, &msg)
	return &msg, err
}

// Encode encodes provided chat Message in binary format
func (m *Message) Encode() ([]byte, error) {
	return msgpack.Marshal(m)
}

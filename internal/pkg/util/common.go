package util

import (
	"github.com/vmihailenco/msgpack"
	"log"
)


// Message represents chat message
type Message struct {
	Meta     map[string]string `json:"meta"`
	Time     int64             `json:"time"`
	Seq      uint64            `json:"seq"`
	Text     string            `json:"text"`
	FromUID  string            `json:"from_uid"`
	FromName string            `json:"from_name"`
	MsgType int 		   `json:"msg_type,omitempty"`
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

func NextId(){

}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


//
//func NewNats() nats.Client {
//	clusterId := viper.GetString("nats.cluster_id")
//	clientId := viper.GetString("nats.client_id")
//	url := viper.GetString("nats.url")
//	mq, err := nats.New(clusterId, clientId, url)
//	CheckErr(err)
//
//	return mq
//}


package nats

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"sync"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

var Info MQInfo

type Client interface {
	Init(mqInfo *MQInfo)
	SubscribeQueue(string,func(uint64, []byte)) (io.Closer, error)
	SubscribeSeq(string, string, uint64, func(uint64, []byte)) (io.Closer, error)
	SubscribeTimestamp(string, string, time.Time, func(uint64, []byte)) (io.Closer, error)
	Send(string, []byte) error
}

type MQInfo struct {
	Mu *sync.RWMutex
	stan stan.Conn
	ClusterId string  `yaml:"cluster_id"`
	ClientId string `yaml:"client_id"`
	Url string `yaml:"url"`
}

// client represents NATS client
type client struct {
	cn stan.Conn
}

func NewNats() Client {
	return &client{}
}

func (mq *client) Init(mqInfo *MQInfo){
	Info = MQInfo{
		Mu: &sync.RWMutex{},
		ClusterId:mqInfo.ClusterId,
		ClientId:mqInfo.ClientId,
		Url:mqInfo.Url,
	}
	conn, err := stan.Connect(mqInfo.ClusterId, mqInfo.ClientId, stan.NatsURL(mqInfo.Url))
	if err != nil {
		panic(err)
	}


	Info.stan = conn
	mq.cn = Info.stan
}

// New initializes a connection to NATS server
func New(clusterID, clientID, url string) (Client, error) {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		return nil, fmt.Errorf("error connecting to NATS: %v", err)
	}
	return &client{cn: conn}, nil
}

// SubscribeQueue subscribers to a message queue
func (c *client) SubscribeQueue(subj string, f func(uint64, []byte)) (io.Closer, error) {
	return c.cn.QueueSubscribe(
		subj,
		"ingest",
		func(m *stan.Msg) {
			f(m.Sequence, m.Data)
		},
		stan.SetManualAckMode(),
	)
}

// SubscribeSeq subscribers to a message queue from received sequence
func (c *client) SubscribeSeq(id string, nick string, start uint64, f func(uint64, []byte)) (io.Closer, error) {
	logs.Info("SubscribeSeq 1")
	var duration_Seconds time.Duration = 5 * time.Second
	return c.cn.Subscribe(
		id,
		func(m *stan.Msg) {
			defer m.Ack()
			logs.Info("SubscribeSeq 2")
			f(m.Sequence, m.Data)
		},
		stan.StartAtSequence(start),
		stan.SetManualAckMode(),
		stan.AckWait(duration_Seconds),
	)
}

// SubscribeTimestamp subscribers to a message queue from received time.Time
func (c *client) SubscribeTimestamp(id string, nick string, t time.Time, f func(uint64, []byte)) (io.Closer, error) {
	return c.cn.Subscribe(
		id,
		func(m *stan.Msg) {
			f(m.Sequence, m.Data)
		},
		stan.StartAtTime(t),
		stan.SetManualAckMode(),
	)
}

// Send publishes new message
func (c *client) Send(id string, msg []byte) error {
	return c.cn.Publish(id, msg)
}


func GetSelfMQ()  Client {
	Info.Mu.RLock()
	defer Info.Mu.RUnlock()
	return &client{cn:Info.stan}
}

func DefaultInfo() MQInfo{

	info := MQInfo{
		Mu: &sync.RWMutex{},
		ClusterId:"test-cluster",
		ClientId:"test-client",
		Url:"nats://127.0.0.1:4222",
	}
	conn, _ := stan.Connect(info.ClusterId, info.ClientId, stan.NatsURL(info.Url))


	info.stan = conn

	return info
}

package redis

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strconv"

	"github.com/vmihailenco/msgpack"
	"../util"
	"github.com/go-redis/redis"
)

const (
	chanListKey             = "channel.list"
	historyPrefix           = "history"
	chatPrefix              = "chat"
	chatLastSeqPrefix       = "last_seq"
	chatClientLastSeqPrefix = "client.last_seq"

	maxHistorySize int64 = 1000
)

type Client interface {
	Get(string) (string, error)
	GetRecent(string, int64) ([]util.Message, uint64, error)
	AppendMessage(string, *util.Message) error
	updateChannelSeq(string, uint64)
	UpdateLastClientSeq(string, string, uint64)
	GetUnreadCount(string, string) uint64
	ListChannels() ([]string, error)
	Save(interface{}) error
}

// Client represents Redis client
type client struct {
	cl *redis.Client
}

func NewRedis() Client{
	addr := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	port := viper.GetInt("redis.port")
	store, err := New(addr, password, port)
	util.CheckErr(err)

	return store
}

// New instantiates new Redis client
func New(addr, pass string, port int) (Client, error) {
	opts := redis.Options{
		Addr: addr + ":" + strconv.Itoa(port),
	}
	if pass != "" {
		opts.Password = pass
	}

	rClient := redis.NewClient(&opts)

	_, err := rClient.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Redis Addr %v, Port %v Reason %v", addr, port, err)
	}
	return &client{cl: rClient}, nil
}

// Get retrieves chat from Client,and need call DecodeChat
func (s *client) Get(id string) (string, error) {
	val, err := s.cl.Get(chatID(id)).Result()
	if err != nil {
		return "", err
	}

	return val,err
}

// GetRecent returns list of recent messages, and sequence until last message
func (s *client) GetRecent(id string, n int64) ([]util.Message, uint64, error) {
	cmd := s.cl.LRange(chatHistoryID(id), -n, -1)
	if cmd.Err() != nil {
		return nil, 0, cmd.Err()
	}

	data, err := cmd.Result()
	if err != nil {
		return nil, 0, err
	}

	if data == nil || len(data) == 0 {
		return nil, 0, nil
	}

	var seq uint64
	msgs := make([]util.Message, len(data))

	for i, m := range data {
		msg, err := util.DecodeMsg([]byte(m))
		if err != nil {
			msg.Text = "message unavailable!"
		} else {
			seq = msgs[i].Seq
		}
	}

	return msgs, (seq + 1), nil
}

// AppendMessage adds new message
func (s *client) AppendMessage(id string, m *util.Message) error {
	data, err := m.Encode()
	if err != nil {
		data, _ = msgpack.Marshal([]byte(`{"text":"message unavailable, unable to encode","from":"util/client"}`))
	}

	key := chatHistoryID(id)

	if err := s.cl.RPush(key, data).Err(); err != nil {
		return err
	}

	s.updateChannelSeq(id, m.Seq)

	return s.cl.LTrim(key, -maxHistorySize, -1).Err()
}

func (s *client) updateChannelSeq(id string, seq uint64) {
	var currSeq int64

	val, err := s.cl.Get(chatLastSeqID(id)).Result()
	if err != nil {
		if err != redis.Nil {
			return
		}
		val = "0"
	}

	currSeq, _ = strconv.ParseInt(val, 10, 64)

	if uint64(currSeq) >= seq {
		return
	}

	s.cl.Set(chatLastSeqID(id), seq, 0)
}

// UpdateLastClientSeq updates client's last seen message
func (s *client) UpdateLastClientSeq(uid string, id string, seq uint64) {
	var currSeq int64

	val, err := s.cl.Get(chatClientLastSeqID(uid, id)).Result()
	if err != nil {
		if err != redis.Nil {
			return
		}
		val = "0"
	}

	currSeq, _ = strconv.ParseInt(val, 10, 64)

	if uint64(currSeq) >= seq {
		return
	}

	s.cl.Set(chatClientLastSeqID(uid, id), seq, 0)
}

// GetUnreadCount returns number of unread messages
func (s *client) GetUnreadCount(uid string, id string) uint64 {
	val, err := s.cl.Get(chatClientLastSeqID(uid, id)).Result()
	if err != nil {
		if err != redis.Nil {
			return 0
		}
		val = "0"
	}

	useq, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}

	val, err = s.cl.Get(chatLastSeqID(id)).Result()
	if err != nil {
		return 0
	}

	cseq, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}

	delta := cseq - useq

	if delta <= 0 {
		return 0
	}

	return uint64(delta)
}

// Save saves new chat
func (s *client) Save(c interface{}) error {
	ctName := ""
	ctSecret := ""
	var data []byte





	name := reflect.ValueOf(c)
	if name.Kind() != reflect.Struct {
		fmt.Println("不是 Struct")
		return nil
	}
	for i:=0; i< name.NumField();i++{
		k:= name.Type().Field(i)
		fmt.Println(k.Name)
		switch k.Name {
		case "Name":
			ctName = k.Name
		case "Secret":
			ctSecret = k.Name

		}
	}

	cc := reflect.TypeOf(c)
	for it := 0; it < cc.NumMethod(); it++ {
		name := cc.Method(it).Name

		switch name {
		case "Encode":
			f, _ := cc.MethodByName(name)
			data = f.Func.Call([]reflect.Value{})[0].Bytes()
			err := f.Func.Call([]reflect.Value{})[1].IsNil()
			if !err  {
				return nil
			}
		}

	}

	pipe := s.cl.TxPipeline()
	pipe.Set(chatID(ctName), data, 0)

	// Save only public channels
	if ctSecret == "" {
		pipe.SAdd(chanListKey, ctName)
	}

	_, err := pipe.Exec()


	return err
}

// ListChannels returns list of all channels
func (s *client) ListChannels() ([]string, error) {
	return s.cl.SMembers(chanListKey).Result()
}

func chatID(id string) string {
	return fmt.Sprintf("%s.%s", chatPrefix, id)
}

func chatHistoryID(id string) string {
	return fmt.Sprintf("%s.%s.%s", historyPrefix, chatPrefix, id)
}

func chatLastSeqID(id string) string {
	return fmt.Sprintf("%s.%s.%s", chatLastSeqPrefix, chatPrefix, id)
}

func chatClientLastSeqID(uid, id string) string {
	return fmt.Sprintf("%s.%s.%s", chatClientLastSeqPrefix, uid, id)
}



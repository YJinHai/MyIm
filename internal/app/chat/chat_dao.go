package chat

import (
	"github.com/YJinHai/MyIm/internal/app/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/rs/xid"
	"github.com/vmihailenco/msgpack"
)

// Limit represents limit type
type Limit int

// Limit constants
const (
	DisplayNameLimit Limit = iota + 1
	UIDLimit
	SecretLimit
	ChanLimit
	ChanSecretLimit
)

// Chat represents private or channel chat
type Chat struct {
	Name    string           `json:"name"`
	Secret  string           `json:"secret"`
	Members map[int]*models.ImUser `json:"members"`
}

// Chat errors
var (
	errAlreadyRegistered = errors.New("chat: uid already registered in this chat")
	errNotRegistered     = errors.New("chat: not a member of this channel")
	errInvalidSecret     = errors.New("chat: invalid secret")
)

// NewChannel creates new channel chat
func NewChannel(name string, private bool) *Chat {
	ch := Chat{
		Name:    name,
		Members: make(map[int]*models.ImUser),
	}

	if private {
		ch.Secret = newSecret()
	}

	return &ch
}

// Register registers user with a chat and returns secret which should
// be stored on the client side, and used for subsequent join requests
func (c *Chat) Register(u *models.ImUser) (string, error) {
	if _, ok := c.Members[u.Uid]; ok {
		return "", errAlreadyRegistered
	}
	if u.Secret == "" {
		u.Secret = newSecret()
	}
	c.Members[u.Uid] = u
	return u.Secret, nil
}

// Join attempts to join user to chat
func (c *Chat) Join(uid int, secret string) (*models.ImUser, error) {
	u, ok := c.Members[uid]
	if !ok {
		return nil, errNotRegistered
	}
	if u.Secret != secret {
		return nil, errInvalidSecret
	}
	u.Secret = ""
	logs.Info("Join:",c.Members[uid])
	return u, nil
}

// Leave removes user from channel
func (c *Chat) Leave(uid int) {
	delete(c.Members, uid)
}

// ListMembers returns list of members associated to a chat
func (c *Chat) ListMembers() []*models.ImUser {
	if len(c.Members) < 1 {
		return nil
	}
	var members []*models.ImUser
	for _, u := range c.Members {
		u.Secret = ""
		members = append(members, u)
	}
	return members
}

// Encode encodes provided chat in binary format
func (c *Chat) Encode() ([]byte, error) {
	return msgpack.Marshal(c)
}


func newSecret() string {
	return xid.New().String()
}

// DecodeChat tries to decode binary formatted message in b to Message
func DecodeChat(b string) (*Chat, error) {
	var c Chat
	if err := msgpack.Unmarshal([]byte(b), &c); err != nil {
		return nil, fmt.Errorf("client: unable to unmarshal chat: %v", err)
	}
	return &c, nil
}


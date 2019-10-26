package chat

type RegisterReq struct {
	Uid           int `json:"uid"`
	Username   string `json:"username"`
	Email         string `json:"email"`
	Secret        string `json:"secret"`
	Channel       string `json:"channel"`
	ChannelSecret string `json:"channel_secret"`
}

type CreateReq struct {
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

type C2CSendRequest struct {
	From      int64 `json:"from"`
	To int64   `json:"to"`
	Content      string `json:"content"`
}

// 推送给接收者的协议
type C2CPushRequest struct {
	MsgId int64 `json:"msg_id"`
	From      int64 `json:"from"`
	Content      string `json:"content"`
}



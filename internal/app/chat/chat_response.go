package chat

type RegisterResp struct {
	Secret string `json:"secret"`
}

type C2CSendResponse struct {
	MsgId int64 `json:"msg_id"`
}

type C2CPushResponse struct {
	MsgId int64 `json:"msg_id"`  // 消息id，服务器收到这个id可以去置位这个消息已读
}



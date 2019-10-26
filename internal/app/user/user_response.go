package user

// swagger:model
type UpdateResponse struct {
	//用户id
	Uid	int `json:"uid"`
	Username string `json:"username"`
}


// swagger:model
type InfoResponse struct {
	//用户id
	Uid	int `json:"uid"`
	// 昵称
	Username	string `json:"username"`
	Email	string `json:"email"`
}

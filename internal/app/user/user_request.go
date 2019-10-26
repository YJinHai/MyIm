package user

// swagger:model
type UpdateRequest struct {
	// the id for this user
	// required: true
	// max length: 64
	Uid  int `json:"uid"`
	Openid   string `json:"openid"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	Email    string`json:"email"`
	Password  string `json:"password"`
}

type RegisterRequest struct {
	Email    string`json:"email"`
	Nickname string `json:"nickname"`
	Password  string `json:"password"`
}

// swagger:model
type InfoRequest struct {
	// 用户ID
	// required: true
	Uid int `json:"uid"`
}


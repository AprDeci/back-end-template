package models

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LogoutReq struct {
	Token string `json:"token"`
}

type LogoutRes struct {
	Msg string `json:"msg"`
}

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRes struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

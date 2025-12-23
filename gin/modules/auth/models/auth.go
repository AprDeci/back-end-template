package models

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

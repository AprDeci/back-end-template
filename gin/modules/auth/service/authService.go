package service

import "gin-template/modules/auth/models"

func Login(loginReq *models.LoginReq) (LoginRes *models.LoginRes, err error) {
	LoginRes = &models.LoginRes{
		ID:       1,
		Username: loginReq.Username,
		Token:    "token",
	}
	//TODO:生成Token

	return LoginRes, nil
}

func Logout(logoutReq *models.LogoutReq) (LogoutRes *models.LogoutRes, err error) {
	return &models.LogoutRes{Msg: "logout success"}, nil
}

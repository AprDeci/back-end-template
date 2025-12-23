package service

import "gin-template/modules/auth/models"

func Login(loginReq *models.LoginReq) (LoginRes *models.LoginRes, err error) {
	LoginRes = &models.LoginRes{
		ID:       1,
		Username: loginReq.Username,
		Token:    "token",
	}
	return LoginRes, nil
}

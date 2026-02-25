package service

import (
	"errors"
	"gin-template/global"
	"gin-template/modules/auth/models"
	userModel "gin-template/modules/users/models"
	"gin-template/utils"

	"golang.org/x/crypto/bcrypt"
)

func Login(loginReq *models.LoginReq) (LoginRes *models.LoginRes, err error) {
	LoginRes = &models.LoginRes{
		ID:       1,
		Username: loginReq.Username,
	}
	// 查询用户信息
	var user userModel.User
	err = global.GVA_DB.Model(&userModel.User{}).Where("username = ?", loginReq.Username).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	// 设置返回值
	LoginRes.ID = int32(user.ID)
	LoginRes.Username = user.Username
	// 生成Token
	LoginRes.Token, err = utils.GenerateTokenWithUserInfo(LoginRes.ID, LoginRes.Username, "user")
	if err != nil {
		return nil, err
	}

	return LoginRes, nil
}

func Logout(logoutReq *models.LogoutReq) (LogoutRes *models.LogoutRes, err error) {
	// TODO: 实现登出逻辑，比如将token加入黑名单
	return &models.LogoutRes{Msg: "logout success"}, nil
}


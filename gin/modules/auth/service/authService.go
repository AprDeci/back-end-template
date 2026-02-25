package service

import (
	"errors"
	"gin-template/global"
	"gin-template/modules/auth/models"
	"gin-template/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authUser struct {
	ID       uint   `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (authUser) TableName() string {
	return "users"
}

func Login(loginReq *models.LoginReq) (LoginRes *models.LoginRes, err error) {
	LoginRes = &models.LoginRes{
		ID:       1,
		Username: loginReq.Username,
	}
	if loginReq.Username == "" || loginReq.Password == "" {
		return nil, errors.New("username and password are required")
	}

	var user authUser
	err = global.GVA_DB.Select("id", "username", "password").Where("username = ?", loginReq.Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	LoginRes.ID = int32(user.ID)
	LoginRes.Username = user.Username
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

func Register(registerReq *models.RegisterReq) (RegisterRes *models.RegisterRes, err error) {
	if registerReq.Username == "" || registerReq.Password == "" {
		return nil, errors.New("username and password are required")
	}

	var existing authUser
	err = global.GVA_DB.Select("id").Where("username = ?", registerReq.Username).First(&existing).Error
	if err == nil {
		return nil, errors.New("username already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := authUser{
		Username: registerReq.Username,
		Password: string(hash),
	}

	err = global.GVA_DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &models.RegisterRes{
		ID:       newUser.ID,
		Username: newUser.Username,
	}, nil
}

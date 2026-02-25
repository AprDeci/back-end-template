package api

import (
	"gin-template/core/response"
	"gin-template/modules/auth/models"
	"gin-template/modules/auth/service"

	"github.com/gin-gonic/gin"
)

// @Summary login
// @Description Authenticate user and return token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginReq body models.LoginReq true "Login request"
// @Success 200 {object} models.LoginRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var loginReq models.LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		response.Fail(c, response.ParamError, err.Error())
		return
	}
	loginRes, err := service.Login(&loginReq)

	if err != nil {
		response.Fail(c, response.ServerError, err.Error())
		return
	}
	response.Success(c, loginRes, "login success")
}

// @Summary logout
// @Description Revoke user token
// @Tags auth
// @Accept json
// @Produce json
// @Param logoutReq body models.LogoutReq true "Logout request"
// @Success 200 {object} models.LogoutRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/auth/logout [post]
func Logout(c *gin.Context) {
	var logoutReq models.LogoutReq
	err := c.ShouldBindJSON(&logoutReq)
	if err != nil {
		response.Fail(c, response.ParamError, err.Error())
		return
	}
	logoutRes, err := service.Logout(&logoutReq)
	if err != nil {
		response.Fail(c, response.ServerError, err.Error())
		return
	}
	response.Success(c, logoutRes, "logout success")
}

// @Summary register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param registerReq body models.RegisterReq true "Register request"
// @Success 200 {object} models.RegisterRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var registerReq models.RegisterReq
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		response.Fail(c, response.ParamError, err.Error())
		return
	}

	registerRes, err := service.Register(&registerReq)
	if err != nil {
		response.Fail(c, response.ServerError, err.Error())
		return
	}

	response.Success(c, registerRes, "register success")
}

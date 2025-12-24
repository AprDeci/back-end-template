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
	err := c.ShouldBindBodyWithJSON(&loginReq)
	if err != nil {
		response.Fail(c, 400, err.Error())
	}
	loginRes, err := service.Login(&loginReq)

	if err != nil {
		response.Fail(c, 500, err.Error())
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
	err := c.ShouldBindBodyWithJSON(&logoutReq)
	if err != nil {
		response.Fail(c, 400, err.Error())
	}
	logoutRes, err := service.Logout(&logoutReq)
	if err != nil {
		response.Fail(c, 500, err.Error())
	}
	response.Success(c, logoutRes, "logout success")
}

package api

import (
	"gin-template/modules/auth/models"
	"gin-template/modules/auth/service"

	"github.com/gin-gonic/gin"
)

// @Summary User login
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
		c.JSON(400, err.Error())
	}
	loginRes, err := service.Login(&loginReq)

	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, loginRes)
}

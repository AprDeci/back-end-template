package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(c *gin.Context, data interface{}, msg string) {
	JSON(c, CodeSuccess, data, msg)
}

func Fail(c *gin.Context, code int, msg string) {
	JSON(c, code, nil, msg)
}

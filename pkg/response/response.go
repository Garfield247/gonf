package response

import (
	"net/http"

	"github.com/Garfield247/gonf.git/pkg/e"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessWithResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: e.StatusOK,
		Msg:  e.StatusText(e.StatusOK),
		Data: data,
	})
}

func FailWithResponse(c *gin.Context, code int, msg string) {
	var message string
	if msg != "" {
		message = msg
	} else {
		message = e.StatusText(code)
	}
	c.JSON(http.StatusInternalServerError, Response{
		Code: code,
		Msg:  message,
		Data: map[string]string{},
	})
}

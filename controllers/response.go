package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code" : 10001,  //程序中的错误码
	"msg" : xx,  //提示信息
	"data" : {} //数据
}
*/


// 定义程序中响应的内容
type ResponseData struct{
	Code ResCode `json:"code" `
	Msg interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
// ResponseError 返回错误响应
func ResponseError(c *gin.Context, code ResCode){
	rd := &ResponseData{
		code,
		code.Msg(),
		nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseWithMsg(c *gin.Context, code ResCode, msg interface{}){
	rd := &ResponseData{
		code,
		msg,
		nil,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseSuccess 返回正常响应
func ResponseSuccess(c *gin.Context, data interface{}){
	rd := &ResponseData{
		CodeSuccess,
		CodeSuccess.Msg(),
		data,
	}
	c.JSON(http.StatusOK, rd)
}
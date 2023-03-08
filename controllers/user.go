package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)
// controller 为服务的入口，负责处理路由，参数校验，请求转发
func SignUpHandler(c *gin.Context){
//	1. 获取参数和参数校验
	var params models.ParamSignUp
	if err:=c.ShouldBindJSON(&params);err!=nil {
		//记录日志
		zap.L().Error("signup with invalid params", zap.Error(err))
		//返回错误，判断errs是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok{
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//	请求参数有误，直接返回
		ResponseWithMsg(c, CodeInvalidPassword, err.Translate(trans))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Translate(trans),
		//})
		return
		}
		fmt.Println(params)
//	2. 业务处理
	if err := logic.SignUp(&params);err != nil{
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败",
		//})
		if errors.Is(err, mysql.ErrorUserExist){
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
//	3. 返回响应
	ResponseSuccess(c,nil)
//	c.JSON(http.StatusOK, gin.H{
//		"status": "ok",
//	})
//
}

func LoginHandler(c *gin.Context){
//	1. 获取参数校验及参数校验
	params := new(models.ParamLogin)
	if err := c.ShouldBindJSON(params);err != nil{
	//	记录日志，
		zap.L().Error("login with invalid params", zap.Error(err))
	//	返回错误，判断error是不是validator.validation错误
		err, ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseError(c, CodeInvalidParam)
		}

		ResponseWithMsg(c, CodeInvalidParam,err.Translate(trans))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Translate(trans),   //翻译错误
		//})
		//return
	}
//	2. 业务处理
	user, err := logic.Login(params)
	if err != nil{
		//记录日志
		zap.L().Error("logic.Login failed to login", zap.String("username", params.Username),zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist){
			ResponseError(c, CodeUserNotExist)
			return

		}
		ResponseError(c, CodeInvalidPassword)

		return
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或密码错误",
		//})

	}
//	3. 返回响应
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "登录成功",
//	})
	ResponseSuccess(c,gin.H{
		"userId": user.UserID,
		"username": user.Username,
		"token": user.Token,
	})


}

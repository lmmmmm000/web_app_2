package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web_app/dao/mysql"
)


const CtxUserIDKey = "userId"

// GetCurrentUser 获取当前登录的用户ID

func GetCurrentUserId(c *gin.Context)(userId int64, err error){
	uid, ok := c.Get(CtxUserIDKey)
	if !ok{
		err = mysql.ErrorUserNotLogin
		return
	}
	userId, ok =  uid.(int64)
	if !ok{
		err = mysql.ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context)(int64, int64){
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var(
		page int64
		size int64
		err error
	)
	page, err = strconv.ParseInt(pageStr, 10,64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10,64)
	if err != nil {
		size = 10
	}
	return page,size
}
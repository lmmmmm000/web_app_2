package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

// CommunityHandler 查询到所有社区
// 社区相关代码
// controller 处理请求参数和路由跳转
func CommunityHandler(c *gin.Context){

//	1. 查询到所有社区(community_id, community_name)以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil{
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)  //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)

}

// CommunityDetailHandler 根据ID查询社区分类详情
func CommunityDetailHandler(c *gin.Context){

	//	获取社区ID
	idStr := c.Param("id")
	//参数校验，查看是不是int型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil{
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil{
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)  //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)

}


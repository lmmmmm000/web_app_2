package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

//投票



func PostVoteController(c *gin.Context){
//	1. 参数检验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p);err != nil {
		errs, ok := err.(validator.ValidationErrors)  //类型断言，传的值有可能没有触发到validator规则
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := errs.Translate(trans)
		ResponseWithMsg(c, CodeInvalidParam, errData)
		return
	}

	// 获取当前请求用户ID
	userId, err := GetCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// PostVote 具体投票的业务逻辑
	if err := logic.VoteForPost(userId, p);err != nil{
		zap.L().Error("logic.VoteForPost", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

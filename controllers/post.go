package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)



// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context){

	// 1. 获取参数及参数校验
	p := new(models.Post)
	//validator --> binding tag
	if err := c.ShouldBindJSON(p); err != nil{
		zap.L().Debug("", zap.Any("err",err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 从c获取当前发请求的用户ID
	userId, err := GetCurrentUserId(c)
	if err!= nil{
		ResponseError(c, CodeNeedLogin)
		return
	}


	p.AuthorId = userId //作者ID
	// 3. 创建帖子
	if err := logic.CreatePost(p);err != nil{
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 4. 返回响应
	ResponseSuccess(c, nil)
}

// CreatePostDetailHandler 获取帖子详情
func CreatePostDetailHandler(c *gin.Context){
//	1. 从URL获取帖子id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil{
		zap.L().Error("get post detail with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

//	2. 根据id取出帖子数据(查数据库)
	data, err := logic.GetPostById(pid)
	if err != nil{
		zap.L().Error("logic.GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return

	//	38821059427831808
	}

//	3. 返回响应
	ResponseSuccess(c, data)
}

// CreatePostListHandler 获取帖子分页展示
func CreatePostListHandler(c *gin.Context){

//	获取分页参数
	page, size := GetPageInfo(c)


//	获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil{
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

//	返回响应

	ResponseSuccess(c, data)

}


// CreatePostListHandler2 获取帖子分页展示
// 根据前端传来的参数动态获取帖子列表
// 按分数、按创建时间、排序
//1.获取参数
//2.去redis查询ID列表
//3.根据id去数据库查询帖子详细信息

// CreatePostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]

func CreatePostListHandler2(c *gin.Context){
	//GET请求参数(querystring) ：/api/v1/posts2?page=1&size=10&order=time
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
		Order: models.OrderTime,
	}
	if err :=c.ShouldBindQuery(p); err != nil{
		zap.L().Error("CreatePostListHandler2 with invalid param ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//gin框架里面用反射机制把url的参数取出来
	//如果请求中携带的是JSON格式的数据才用shouldbindjson
	//	获取分页参数
	//page, size := GetPageInfo(c)    


	//	获取数据
	data, err := logic.GetPostListNew(p) //更新: 合二为一
	if err != nil{
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//	返回响应

	ResponseSuccess(c, data)

}

// 根据社区id查询帖子列表

func GetCommunityPostListHandler (c *gin.Context) {
	//GET请求参数(querystring) ：/api/v1/posts2?page=1&size=10&order=time
	//p := &models.ParamCommunityPostList{
	//	ParamPostList: models.ParamPostList{Page: 1, Size: 10, Order: models.OrderScore},
	//}
	//if err := c.ShouldBindQuery(p); err != nil {
	//	zap.L().Error("CreatePostListHandler2 with invalid param ", zap.Error(err))
	//	ResponseError(c, CodeInvalidParam)
	//	return
	//}
	////gin框架里面用反射机制把url的参数取出来
	////如果请求中携带的是JSON格式的数据才用shouldbindjson
	////	获取分页参数
	////page, size := GetPageInfo(c)
	//
	////	获取数据
	//data, err := logic.GetCommunityPostList2(p)
	//if err != nil {
	//	zap.L().Error("logic.GetPostList failed", zap.Error(err))
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	//ResponseSuccess(c, data)
}

package models


const(
	OrderTime = "time"
	OrderScore = "score"
)


//定义请求参数的结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct{
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct{
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct{
	//	userId
	PostId string `json:"post_id" binding:"required"` //帖子ID
	Direction int8 `json:"direction,string" binding:"oneof=-1 0 1" ` //赞成票 (1) 反对票(-1)
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct{
	CommunityID int64 `json:"community_id" form:"community_id"`//可以为空
	Page int64 `json:"page" form:"page"`
	Size int64	`json:"size" form:"size"`
	Order string `json:"order" form:"order" example:"score"`
}

// ParamCommunityPostList 按社区获取帖子列表query string参数
type ParamCommunityPostList struct{
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
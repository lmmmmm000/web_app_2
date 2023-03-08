package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"

	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post)(err error){
	//雪花算法生成ID
	p.ID = snowflake.GetId()
//	保存到数据库
    err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityId)
	return
}

// GetPostById 根据帖子id查询帖子详情
func GetPostById(pid int64)(data *models.ApiPostDetail, err error){
	//查询并组合接口用到的数据

	post, err :=  mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
//	根据帖子中的作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed",
			zap.Int64("post.AuthorId", post.AuthorId),
			zap.Error(err))
		return
	}
//	根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed",
			zap.Int64("post.CommunityId", post.CommunityId),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		user.Username,
		post,
		community,
	}

	return
}


// GetPostList 获取帖子列表
func GetPostList(page, size int64)(data []*models.ApiPostDetail, err error){
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail,0, len(posts))
	for _, post := range posts{
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed",
				zap.Int64("post.AuthorId",post.AuthorId),
				zap.Error(err))
			continue
		}
		community, err:= mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById",
				zap.Any("post.CommunityId", post.CommunityId),
				zap.Error(err))
			continue
		}
		postsDetail := &models.ApiPostDetail{
			user.Username,
			post,
			community,
		}
		data = append(data, postsDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList)(data []*models.ApiPostDetail2, err error){
//	2. 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil{
		return
	}
	if len(ids) == 0{
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
//	3. 根据id去数据库查询帖子的详细信息
	// 返回的数据还有按照给定的数据
	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("mysql.GetPostByIDs", zap.Any("posts", posts))

	//提前查询好每个帖子的投票数
	voteData, err := redis.GetPostData(ids)
	if err != nil {
		return
	}

	data =  make([]*models.ApiPostDetail2, 0, len(posts))
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil{
			zap.L().Error("mysql.GetUserById",  
				zap.Int64("post.AuthorId", post.AuthorId),
				zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailById(user.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById",
				zap.Int64("user.UserID", user.UserID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail2{
			user.Username,
			voteData[idx],
			post,
			community,
		}
		data = append(data, postDetail)
	}
	return

}

func GetCommunityPostList2(p *models.ParamPostList)(data []*models.ApiPostDetail2, err error){
	//	2. 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil{
		return
	}
	if len(ids) == 0{
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	//	3. 根据id去数据库查询帖子的详细信息
	// 返回的数据还有按照给定的数据
	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("mysql.GetPostByIDs", zap.Any("posts", posts))

	//提前查询好每个帖子的投票数
	voteData, err := redis.GetPostData(ids)
	if err != nil {
		return
	}

	data =  make([]*models.ApiPostDetail2, 0, len(posts))
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil{
			zap.L().Error("mysql.GetUserById",
				zap.Int64("post.AuthorId", post.AuthorId),
				zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailById(user.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById",
				zap.Int64("user.UserID", user.UserID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail2{
			user.Username,
			voteData[idx],
			post,
			community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListNew(p *models.ParamPostList)(data []*models.ApiPostDetail2, err error){
	if p.CommunityID == 0{
		data, err = GetPostList2(p)
	}else {
		data, err = GetCommunityPostList2(p)

}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
return
}
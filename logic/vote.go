package logic

import (
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

// 投票功能
// 1. 用户投票的数据
//	基于用户投票的相关算法: http://www.ruanyifeng.com/blog/2012/02/ranking_algorithm_hacker_news.html




// Vote 使用简化版的投票算法
// 投一票就加432分 一天86400秒 / 200 -> 需要200张赞成票可以给帖子续一天 -> redis 实战




// VoteForPost 为帖子投票的函数
func VoteForPost(userId int64, p *models.ParamVoteData)(err error){
	zap.L().Debug("VoteForPost",zap.Int64("userId", userId),
		zap.String("postId", p.PostId),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostId , float64(p.Direction))

}

package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"strconv"

	"math"
	"time"
)
/* 投票的几种情况:
direction = 1, 用两种情况：
	1. 之前没有投过票，现在投赞成票 --> 更新分数和投票记录 差值的绝对值：1 +432
	2. 之前投反对票，现在改投赞成票 --> 更新分数和投票记录 差值的绝对值：2 +432*2
direction = 0, 用两种情况：

	1. 之前投反对票，现在取消投票 --> 更新分数和投票记录 差值的绝对值：1	+432
	2. 之前投过赞成票，现在取消投票 --> 更新分数和投票记录 差值的绝对值：1 -432
direction = -1, 用两种情况：
	1. 之前没有投过票，现在投反对票 --> 更新分数和投票记录 差值的绝对值：1 -432
	2. 之前投赞成票，现在取消投票 --> 更新分数和投票记录 差值的绝对值：2	-432*2

投票的限制：
每个帖子自发表之日起，一个星期之内允许用户投票，超过一个星期就不允许再投票
	1. 到期之后将redis保存的赞成票数及反对票数存储到mysql中
	2. 到期之后删除那个KeyPostVotedZSetPF
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	ScorePerVote = 432 //每一票的分值 432分

)

var(
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated = errors.New("投票重复")
)

func CreatePost(postId, communityID int64)error{
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		float64(time.Now().Unix()),
		postId,
	})
	//帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		float64(time.Now().Unix()),
		postId,
	})
	//把帖子id加到社区set里面
	cKey := GetRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postId)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userId , postId string, value float64)error{

	// 1. 判断投票的限制
	//去redis取帖子发布时间

	//func (cmd *FloatCmd) Val() float64
	PostTime := client.ZScore(GetRedisKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix()) - PostTime > oneWeekInSeconds{
		return ErrVoteTimeExpire
	}

	// 2, 3需要放到一个pipeline里面
	//	2. 更新帖子分数
	// 先查当前用户给当前帖子的投票记录
	oldValue := client.ZScore(GetRedisKey(KeyPostVotedZSetPF+postId), userId).Val()
	if value == oldValue{
		return ErrVoteRepeated
	}
	var op float64
	if value > oldValue{
		op = 1
	}else{
		op = -1}
	diff := math.Abs(oldValue - value) //计算两次投票的差值

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), op*diff*ScorePerVote, postId)

	//	3. 记录用户为该帖子投过票的数据
	if value == 0{
		//如果取消投票就移除改postId
		pipeline.ZRem(GetRedisKey(KeyPostVotedZSetPF+postId))}else{
		pipeline.ZAdd(GetRedisKey(KeyPostVotedZSetPF+postId), redis.Z{
			value,
			userId,
		})
	}
	_, err := pipeline.Exec()
	return err

}


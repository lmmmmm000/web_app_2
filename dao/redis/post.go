package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"web_app/models"
)

func getIDsFromKey(key string, page, size int64)([]string, error){
	start := (page - 1)* size
	end := start + size -1
	//	3. zrevrange 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList)([]string, error){
// 从redis获取id
//	根据用户请求中携带的order参数确定查询的key
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore{
		key = GetRedisKey(KeyPostScoreZSet)
	}

//	2.确定查询的索引起始点
	return getIDsFromKey(key, p.Page, p.Size)

}

// GetPostData 根据ids查询没篇帖子赞成票的数据
func GetPostData(ids []string)(data []int64, err error){
	//data = make([] int64, 0, len(ids))
	//for _, id:=range ids{
	//	key := GetRedisKey(KeyPostVotedZSetPF+id)
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	//使用pipeline一次发送多条命令，减少rtt
	pipeline := client.Pipeline()
	for _, id:=range ids{
		key := GetRedisKey(KeyPostVotedZSetPF+id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders{
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

// GetCommunityPostIDsInOrder 按社区根据ids查询每篇帖子的投赞成票的数据
func GetCommunityPostIDsInOrder(p *models.ParamPostList)([]string, error){
	// 使用zinterstore 把分区的帖子 set与帖子分数的zset生成一个新的zset
	// 针对新的zset 按之前逻辑取数

	orderKey := GetRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore{
		orderKey = GetRedisKey(KeyPostScoreZSet)
	}


	// 利用缓存key减少zinterstore执行的次数
	//社区的key
	cKey:= GetRedisKey(KeyCommunitySetPF+strconv.Itoa(int(p.CommunityID)))
	// 从redis获取id
	//	根据用户请求中携带的order参数确定查询的key
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val()<1 {
	//	不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) // zintstore 计算
		pipeline.Expire(key, 60 * time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}


	//	存在就直接根据key查询ids

	return getIDsFromKey(key, p.Page, p.Size)
}
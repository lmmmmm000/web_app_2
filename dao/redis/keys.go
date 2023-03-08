package redis

// redis key

//redis key 尽量使用命名空间的方式，区分不同的key，方便查询和拆分
const (
	Prefix = "bluebell:"
	KeyPostTimeZSet = "post:time" //zset;帖子及发帖时间为分数
	KeyPostScoreZSet = "post:score" //zset;帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" //zset;记录用户以及投票的类型, 参数是post id(后面带冒号的是不完整的key)

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
	)

func GetRedisKey(key string)string{
	return Prefix + key
}
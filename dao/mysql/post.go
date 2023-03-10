package mysql

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"web_app/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post)(err error){
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?, ?, ?, ?, ?)
		`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorId, p.CommunityId)
	return
}

// GetPostById 根据id查询单个帖子数据
func GetPostById(pid int64)(post *models.Post, err error){
	post = new(models.Post)
	sqlStr := `select 
       		post_id, title, content, author_id, community_id, create_time
			from post
			where post_id = ?
			`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64)(posts []*models.Post, err error){
	posts = make([] *models.Post, 0, 2)
	sqlStr := `select
				post_id, title, content, author_id, community_id, create_time
				from post
				order by create_time
				desc
				limit ?,?
				`
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return

}

// GetPostByIDs 根据给的id列表查询帖子

func GetPostByIDs(ids []string)(postList[] *models.Post, err error){
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in(?)
	order by FIND_IN_SET(post_id, ?)
`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	//rebind 用法
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
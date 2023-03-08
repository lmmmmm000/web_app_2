package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList()([]*models.Community, error){
	// 查找到所有的Community并返回
	return mysql.GetCommunityList()

}


func GetCommunityDetail(id int64)(*models.CommunityDetail, error){
	// 查找到指定ID的Community并返回
	return mysql.GetCommunityDetailById(id)

}
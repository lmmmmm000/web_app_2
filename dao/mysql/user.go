package mysql

import (
 "crypto/md5"
 "database/sql"
 "encoding/hex"
 "web_app/models"
)

// 把每一步数据库操作封装成函数

// 等logicc层根据业务需求调用

// CheckUserExist 检查指定用户名是否存在

var secrect = "www.google.com"

func CheckUserExist(username string) (err error){
 sqlStr := `select count(user_id) from user where username=?`
 var count int
 if err := db.Get(&count, sqlStr, username);err != nil{
  return err
 }
 if count>0 {
  return ErrorUserExist

 }
 return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User)(err error){
 //对密码进行加密
 user.Password = encryptPassword(user.Password)
//	1. 执行数据入库
 sqlStr := `INSERT INTO user(user_id, username, password) values(?,?,?)`
 _, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
 return

}

// encryptPassword 验证密码
func encryptPassword(oldPassword string)string{
 h :=md5.New()
 h.Write([]byte(secrect))
 return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}

func Login(user *models.User)(err error){
 oldPassword := user.Password //用户登录的密码
 sqlStr := `select user_id, username, password from user where username=?`

 err=db.Get(user, sqlStr, user.Username) //这里user本来就是指针变量，不用再传指针
 if err == sql.ErrNoRows{
  return ErrorUserNotExist
 }
 if err !=nil{
  return err
 }
// 判断密码是否正确
password := encryptPassword(oldPassword)
if password != user.Password{
 return ErrorInvalidPassword
}
return

}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64)(user *models.User, err error){
 user = new(models.User)
 sqlStr := `select user_id, username from user where user_id=?`
 db.Get(user, sqlStr, uid)
 return
}
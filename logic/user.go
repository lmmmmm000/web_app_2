package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp)(err error){
	//1.判断用户是否存在
	 if err = mysql.CheckUserExist(p.Username);err != nil{
		//数据库查询出错
		return err
	}

	//2. 生成UID
	userID := snowflake.GetId()
	//构造user实例
	user := &models.User{
		userID,
		p.Username,
		p.Password,
		"",
	}


	//3. 用户密码要加密

	//4. 保存进数据库
	err = mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return
}

func Login(p *models.ParamLogin)(user *models.User,err error){
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//user := &models.ParamLogin{
	//	p.Username,
	//	p.Password,
	//}

	//传递的是指针
	if err :=  mysql.Login(user);err != nil{
		return nil, err
	}
//	生成jwt的token
 	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return

}

package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)



var mySecret = []byte("柳鸣牛逼")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId int64 `json:"user_id"`
	Username string `json:"username"`

	jwt.StandardClaims
}

// GenToken 生成jwt
func GenToken(userId int64, username string)(string, error){
	c := MyClaims{
		userId, //定义字段
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire"))*time.Hour).Unix(), //过期时间
			Issuer: "web_app_2", //过期时间签发人
		},
	}
//	使用指定的签名方法创建签名对象
token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//使指定的secret签名并获得完整的编码后的字符串
return token.SignedString(mySecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var  mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
package auth

import (
	"fmt"
	"time"

	"code.byted.org/zhuchaowen/trae/config"
	"code.byted.org/zhuchaowen/trae/config/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 密钥，用于签名和验证JWT
var jwtKey = []byte("your_secret_key")

// 定义用户凭证结构体
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义声明结构体
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 登录处理函数
func Login(c *gin.Context) {
	var creds Credentials
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		config.Fail(c, constants.ErrBadRequest, "请求参数错误")
		return
	}

	// 从Redis获取用户密码
	ctx := c.Request.Context()
	rdb := config.GetRedisClient()
	expectedPassword, err := rdb.Get(ctx, creds.Username).Result()
	if err != nil {
		config.Fail(c, constants.ErrUnauthorized, "用户名或密码错误")
		return
	}

	if expectedPassword != creds.Password {
		config.Fail(c, constants.ErrUnauthorized, "用户名或密码错误")
		return
	}

	// 设置声明
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		config.Fail(c, constants.ErrInternalServer, "生成令牌失败")
		return
	}

	// 将用户信息存储到Redis
	userInfo := map[string]interface{}{
		"username":   creds.Username,
		"expires_at": expirationTime,
	}
	err = rdb.HSet(ctx, creds.Username, userInfo).Err()
	if err != nil {
		config.Fail(c, constants.ErrInternalServer, "存储用户信息到Redis失败")
		return
	}

	// 返回令牌
	c.SetCookie("token", tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "", false, true)
	config.Success(c, gin.H{"message": "登录成功"})
}

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie中的令牌
		tokenString, err := c.Cookie("token")
		if err != nil {
			config.Fail(c, constants.ErrUnauthorized, "未提供令牌")
			c.Abort()
			return
		}

		claims := &Claims{}

		// 解析令牌
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				config.Fail(c, constants.ErrUnauthorized, "无效的令牌")
				c.Abort()
				return
			}
			config.Fail(c, constants.ErrBadRequest, "令牌解析错误")
			c.Abort()
			return
		}

		if !token.Valid {
			config.Fail(c, constants.ErrUnauthorized, "无效的令牌")
			c.Abort()
			return
		}

		c.Next()
	}
}

// 受保护的路由处理函数
func Protected(c *gin.Context) {
	// 从令牌中获取用户名
	claims := c.MustGet("claims").(*Claims)
	config.Success(c, gin.H{"message": fmt.Sprintf("欢迎, %s!", claims.Username)})
}

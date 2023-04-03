package middleware

import (
	"gin-go-bl/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// jwt加密密钥
var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint `json:"userid"`
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {

	// token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		// 自定义字段
		UserId: user.ID,
		// 标准字段
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	// 使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	// 返回token
	return tokenString, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "token不存在",
			})
			c.Abort()
			return
		}
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token错误或不存在",
			})
			c.Abort()
			return
		}
		// 提取token的有效部分
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := ParseToken(tokenString)
		// 非法token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token错误",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token过期",
			})
			c.Abort()
			return
		}
		// 获取claims中的userId
		userid := claims.UserId
		DB := model.DB
		var user model.User
		DB.Where("id =?", userid).First(&user)
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()

	}
}

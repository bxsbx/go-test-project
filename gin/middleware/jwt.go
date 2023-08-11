package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CustomClaims struct {
	UUID        uuid.UUID `json:"uuid"`
	Origin      int       `json:"origin"`
	ID          uint      `json:"id,omitempty"`
	Username    string    `json:"userName,omitempty"`
	NickName    string    `json:"nickName"`
	AuthorityId string    `json:"authorityId,omitempty"`
	BufferTime  int64     `json:"bufferTime"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey  []byte
	ExpiresTime int64 // 过期时间
	BufferTime  int64 // 缓冲时间
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

var j JWT

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-auth-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-auth-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, TokenInvalid)
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + j.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

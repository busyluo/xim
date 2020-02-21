package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// LoginRequest 登录提供内容
type LoginRequest struct {
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}

// Token 以上第四步返回给客户端的token对象
type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    int       `json:"-"`
}

func Login(c echo.Context) error {
	// 判断何种方式登录，小程序为提供code
	var req = new(LoginRequest) // 输入请求
	if err := c.Bind(req); err != nil {
		return err
	}
	var t *Token
	if req.Username == "username" && req.Password == "password" {
		// 发行token
		t = &Token{
			Token:     newUUID(),
			ExpiresAt: time.Now().Add(time.Hour * 96),
			// 这个userid应该是检索出来的，这里为demo写死。
			UserID: 1,
		}
		//setcc("token:"+t.Token, t, time.Hour*96)
	} else {
		//return ErrAuthFailed
	}
	return c.JSON(http.StatusOK, t)
}

func newUUID() string {

}

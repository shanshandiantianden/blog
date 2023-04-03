package v1

import (
	"gin-go-bl/middleware"
	"gin-go-bl/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InIndex(c *gin.Context) {
	c.HTML(200, "come.html", nil)
}

func Login(c *gin.Context) {
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	//user := c.PostForm("username")
	//password := c.PostForm("password")
	user := data.UserName
	password := data.Password
	code := model.CheckLogin(user, password)
	if user == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码为空",
		})
		return
	}
	if code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码错误!",
		})
		return
	}
	if code == 200 {
		data = model.GetUserId(user)
		token, _ := middleware.ReleaseToken(data)
		c.JSON(200, gin.H{
			"userid": data.ID,
			"user":   user,
			"token":  token,
			"code":   200,
			"msg":    "登陆成功",
		})
	}
}

func AdminLogin(c *gin.Context) {
	var _ model.Admin
	user := c.PostForm("username")
	password := c.PostForm("password")
	code := model.CheckAdminLogin(user, password)
	if user == "" || password == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"code": 404,
			"msg":  "用户名或密码为空",
		})
		return
	}
	if code != 200 {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"code": 404,
			"msg":  "用户名或密码错误!",
		})
		return
	}
	if code == 200 {
		_ = model.GetAdminId(user)
		c.JSON(200, gin.H{
			"user": user,
			"code": 200,
			"msg":  "登陆成功",
		})
	}
}

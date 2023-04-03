package v1

import (
	"fmt"
	"gin-go-bl/middleware"
	"gin-go-bl/model"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	code int
)

func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	msg, code := middleware.Validate(&data)

	if code != utils.SUCCESS {
		c.JSON(200, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(data.UserName)
	if code == 200 {
		model.CreatUser(&data)
	}
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetUser(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, total := model.GetUser(size, num)
	code = utils.SUCCESS
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func UpdateUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	fmt.Println(data)
	code = model.CheckUser(data.UserName)
	if code == 200 {
		model.UpdateUser(id, &data)
	}
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.DeleteUser(id)
	c.JSON(200, gin.H{
		"status":  data,
		"message": utils.GetErrMsg(data),
	})
}

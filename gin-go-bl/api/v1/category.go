package v1

import (
	"gin-go-bl/model"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddCategory(c *gin.Context) {
	user, _ := c.Get("user")
	id := user.(model.User).ID
	var data model.Category
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	if id == 1 {
		c.JSON(200, gin.H{
			"data":    "",
			"status":  0,
			"message": "非管理员用户,创建失败",
		})
		return
	}
	code = model.CheckCategory(data.CaName)
	if code == 200 {
		model.CreatCategory(&data)
	}
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetCategory(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, total := model.GetCategory(size, num)
	code = utils.SUCCESS
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

//func GetCatag(c *gin.Context) {
//	size, _ := strconv.Atoi(c.Param("size"))
//	num, _ := strconv.Atoi(c.Param("num"))
//	if size == 0 {
//		size = -1
//	}
//	if num == 0 {
//		num = -1
//	}
//
//	data, total := model.GetCategory(size, num)
//	code = utils.SUCCESS
//	c.HTML(200, "tags.html", gin.H{
//		"data":    data,
//		"size":    total,
//		"status":  code,
//		"message": utils.GetErrMsg(code),
//	})
//}

func UpdateCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code = model.CheckCategory(data.CaName)
	if code == 200 {
		model.UpdateCategory(id, &data)
	}
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.DeleteCategory(id)
	c.JSON(200, gin.H{
		"status":  data,
		"message": utils.GetErrMsg(data),
	})
}

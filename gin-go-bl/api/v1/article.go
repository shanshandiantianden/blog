package v1

import (
	"fmt"
	"gin-go-bl/model"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AddArticle(c *gin.Context) {
	user, _ := c.Get("user")
	id := user.(model.User).ID
	var data model.Article
	data.UserId = id
	crtime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05.000000000"), time.Local)
	data.CreateTime = crtime
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code = model.CreatArticle(&data)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetAllArticle(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	arts, code, atotal := model.GetAllArticle(size, num)
	//cate, _ := model.GetCategory(100, 0)
	c.JSON(200, gin.H{
		"arts": arts,
		//"cate":    cate,
		"size":    atotal,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetAllArticles(c *gin.Context) {

	arts, code, atotal := model.GetAllArticles()
	//cate, _ := model.GetCategory(100, 0)
	c.JSON(200, gin.H{
		"arts": arts,
		//"cate":    cate,
		"size":    atotal,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetUserArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := model.GetUserArticle(id, size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetMeArticle(c *gin.Context) {
	data, _ := c.Get("user")
	id := data.(model.User).ID
	//personValue := reflect.ValueOf(data)
	//ID := personValue.FieldByName("ID").Uint()
	//id := middleware.Userid
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := model.GetUserArticle(int(id), size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetCaArticle(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	id, _ := strconv.Atoi(c.Param("id"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := model.GetCatArticle(id, size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetArticle(id)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code = model.EditArticle(id, &data)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.DeleteArticle(id)
	c.JSON(200, gin.H{
		"status":  data,
		"message": utils.GetErrMsg(data),
	})
}

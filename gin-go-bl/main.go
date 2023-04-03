package main

import (
	"gin-go-bl/middleware"
	"gin-go-bl/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	////r.LoadHTMLGlob("views/*")
	////r.LoadHTMLGlob("views/**/*")
	//r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"index": "默认",
		})
	})
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Loger())
	router.ApiRouter(r)
	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(404, "error.html", nil)
	//})

	if err := r.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	} // 监听并在 0.0.0.0:8090 上启动服务
}

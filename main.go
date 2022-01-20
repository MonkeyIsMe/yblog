package main

import (
	"net/http"
	"yblog/client"
	"yblog/model"

	"github.com/gin-gonic/gin"
)

func main() {

	client.InitClient()

	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "t")
		action := c.DefaultQuery("action", "t")

		tag := &model.Tag{
			TagName: name,
			TagInfo: action,
		}
		tag.AddTag()

		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	router.Run(":8000")
}

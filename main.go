package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.New(goview.Config{
		Root:      "templates",
		Extension: ".tmpl",
		Master:    "layouts/master",
	})
	r.Static("/static", "./static/")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Snowcraft Official",
		})
	})

	r.Run(":25566")
}

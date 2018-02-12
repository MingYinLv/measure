package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"measure/src/config"
	"measure/src/handle"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var listenPort int64
	config.Initial(&listenPort)
	router := gin.Default()

	router.LoadHTMLFiles("./static/index.html", "./static/list.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{})
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/addMeasure", handle.AddMeasure)
	router.GET("/measure/:id", handle.FindById)
	router.GET("/measures", handle.FindList)
	router.GET("/deleteMeasure/:id", handle.DeleteMeasure)
	router.POST("/editMeasure/:id", handle.EditMeasure)
	router.GET("/downloadMeasure/:id", handle.DownloadMeasure)

	router.Static("/static", "./static/static")
	router.Run(fmt.Sprintf(":%d", listenPort))
}

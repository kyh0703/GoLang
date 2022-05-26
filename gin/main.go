package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/gin/router"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	route := gin.Default()
	// r.Use(Logger())

	route.LoadHTMLGlob("views/*")
	route.GET("/", indexHandler)
	route.POST("/", router.FromHandler)
	route.GET("/ping", router.PingPong)
	route.GET("/some", router.GetSomeJSON)
	route.GET("/test-binding", router.GetBinding)
	route.POST("/test-query", router.BindQuery)
	route.GET("/test-uri/:name/:id", router.BindUri)
	route.Run(":8080")
}

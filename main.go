package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var banyakCall = 0

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", count)
	router.Run(":9050")
}

func count(c *gin.Context) {
	banyakCall++
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"BanyakCall": banyakCall,
	})
}

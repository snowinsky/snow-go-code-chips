package api

import (
	"log"
	"net/http"
)

import "github.com/gin-gonic/gin"

func GinDemo() {

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"code":    "0000",
			"message": "通信成功",
		})
	})
	r.GET("/person/name", func(c *gin.Context) {
		var name = c.Query("name")
		log.Println("/person/name?name=", name)

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"code":    "0000",
			"message": "通信成功",
			"data":    "name=" + name,
		})
	})
	r.Run()

}

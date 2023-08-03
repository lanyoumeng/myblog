package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/data", func(c *gin.Context) {
		data := gin.H{
			"key1": "value1",
			"key2": 123,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")

}

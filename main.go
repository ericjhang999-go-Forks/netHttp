package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":"github.com/ericjhang999",
			"Line Id":"334455667878",
		})
	})
	r.Run(":8080")
}
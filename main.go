package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var person = struct {
		Name string
		Age  int
	}{
		Name: "golang-gin", Age: 8,
	}
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"person": &person,
		})
	})

	r.Run(":8081")
}

package main

import (
	"net/http"
	// 	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/functionality"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html") //   template.

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

	// r.GET("/users", functionality.GetUsers)
	// r.GET("/users", functionality.GetUsersString)
	r.GET("/users", functionality.GetUsersHTML)

	r.GET("/user", functionality.GetUserByID)

	r.GET("/json-demo", functionality.JsonDemo)
	r.GET("/pure-json-demo", functionality.PureJsonDemo)

	r.GET("/users/:userid/*format", functionality.GetUserByRouterParameters)

	r.GET("/should-bind-query", functionality.UserShouldBindParameter)

	r.POST("/create-user", functionality.CreateUserFromJsonData)

	r.Run(":8081")
}

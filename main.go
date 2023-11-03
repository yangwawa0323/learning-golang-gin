package main

import (
	"net/http"
	// 	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/functionality"
	"github.com/yangwawa0323/learning-golang-gin/middlewares"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(middlewares.MyLogger(), gin.Logger(), gin.Recovery())

	// use QPS limiter middleware
	limit := middlewares.NewLimiter()
	r.Use(limit.RateLimit(1, 3))

	// favicon.ico static file
	r.StaticFile("/favicon.ico", "./favicon.ico")

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

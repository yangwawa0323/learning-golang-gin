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
	r.Use(gin.Logger(), gin.Recovery())

	// use QPS limiter middleware
	limit := middlewares.NewLimiter()
	// r.Use(limit.RateLimit(1, 3))

	// favicon.ico static file
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.LoadHTMLGlob("templates/*.html") //   template.

	var person = struct {
		Name string
		Age  int
	}{
		Name: "golang-gin", Age: 8,
	}

	v10 := r.Group("/v10.01")
	{
		v10.GET("/welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"person": &person,
			})
		})
	}

	// r.GET("/users", functionality.GetUsers)
	// r.GET("/users", functionality.GetUsersString)
	v20 := r.Group("/v20.09")
	{
		v20.GET("/users", functionality.GetUsersHTML)

		v20.GET("/user", functionality.GetUserByID)

		v20.GET("/json-demo", functionality.JsonDemo)
	}

	api := r.Group("/api")
	{
		api.GET("/pure-json-demo", functionality.PureJsonDemo)

		api.GET("/users/:userid/*format", functionality.GetUserByRouterParameters)

		api.GET("/should-bind-query", functionality.UserShouldBindParameter)

		api.POST("/create-user", functionality.CreateUserFromJsonData)

	}
	// r.GET("/frontend", functionality.FrontEnd)

	// r.GET("/frontend-v1", functionality.FrontEndV1)
	// r.GET("/frontend-v2", functionality.FrontEndV2)

	r.POST("user-favs", limit.RateLimit(1, 3), functionality.PostUserFavorites)

	//      /admin/summary      /admin/total     /admin/secret
	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "password",
		"john":  "password",
	}))
	{
		admin.GET("/summary", functionality.AdminSummary)
		admin.GET("/total", functionality.AdminTotal)
		admin.GET("/secret", functionality.AdminSecret)
	}

	r.Run(":8081")
}

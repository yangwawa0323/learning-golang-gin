package functionality

import (
	"bytes"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/dao"
	"github.com/yangwawa0323/learning-golang-gin/model"
	"gorm.io/gorm"
)

type User map[string]interface{} // 👎

// reflect
var GetUsers = func(c *gin.Context) {
	var users []map[string]interface{}
	dao.DB.Raw("SELECT * FROM users LIMIT 50").Scan(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": &users,
	})

}

var GetUsersString = func(c *gin.Context) {
	var users []model.User
	dao.DB.Limit(30).Find(&users)

	var output bytes.Buffer
	output.WriteString("<ol class='user-list-item'>")
	for _, user := range users {
		output.WriteString("<li style='color:blue'>" + user.Name + "</li>")
	}
	output.WriteString("<ol>")

	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, output.String())

}

var GetUsersHTML = func(c *gin.Context) {
	var users []model.User
	dao.DB.Limit(30).Find(&users)

	c.HTML(http.StatusOK, "users.html", &users)
}

// 1.   /user?id=10&format=json
// 2.  /user?id=10
// 2.  /user?id=10&format=html
var GetUserByID = func(c *gin.Context) {
	var user model.User

	userid := c.Query("id")
	result := dao.DB.First(&user, userid)
	if result.Error == gorm.ErrRecordNotFound || result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		// IMPORTANT
		return
	}

	// html, json
	format := c.DefaultQuery("format", "html")

	if format == "html" {
		// c.Writer.Header().Set("Content-Type", "text/html")
		// var info bytes.Buffer
		// info.WriteString("<p>" + user.Name + "</p>")
		// info.WriteString("<p>" + *user.Email + "</p>")
		// c.String(http.StatusOK, info.String())

		c.HTML(http.StatusOK, "user.html", &user)
		// IMPORTANT!!!
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})

}

var CreateUserFromJsonData = func(c *gin.Context) {
	var user model.User
	// expect json format --> BindJSON , ShouldBindJSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// IMPORTANT!!
		return
	}

	dao.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "user create succeed.",
		"data":    &user,
	})
}

// /users/:userid/*format/
var GetUserByRouterParameters = func(c *gin.Context) {
	var user model.User

	userid := c.Param("userid")
	result := dao.DB.First(&user, userid)
	if result.Error == gorm.ErrRecordNotFound || result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		// IMPORTANT
		return
	}

	// html, json
	format := c.Param("format")
	log.Printf("format parameter : %q", format)

	if strings.HasPrefix(format, "/html") {
		c.HTML(http.StatusOK, "user.html", &user)

		// IMPORTANT!!!
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

var UserShouldBindParameter = func(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindQuery(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

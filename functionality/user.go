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
	output.WriteString("<ol>")
	for _, user := range users {
		output.WriteString("<li>" + user.Name + "</li>")
	}
	output.WriteString("<ol>")

	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, output.String())

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
		c.Writer.Header().Set("Content-Type", "text/html")
		var info bytes.Buffer
		info.WriteString("<p>" + user.Name + "</p>")
		info.WriteString("<p>" + *user.Email + "</p>")
		c.String(http.StatusOK, info.String())

		// IMPORTANT!!!
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
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
		c.Writer.Header().Set("Content-Type", "text/html")
		var info bytes.Buffer
		info.WriteString("<p>" + user.Name + "</p>")
		info.WriteString("<p>" + *user.Email + "</p>")
		c.String(http.StatusOK, info.String())

		// IMPORTANT!!!
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

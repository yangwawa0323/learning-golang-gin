package functionality

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/dao"
	"github.com/yangwawa0323/learning-golang-gin/model"
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

package functionality

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/dao"
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

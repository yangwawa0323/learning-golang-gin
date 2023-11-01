package functionality

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/learning-golang-gin/model"
)

var Demo = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"person": model.Person,
	})
}

var PostDemo = func(c *gin.Context) {
	c.String(http.StatusOK, "Post demo")
}

var DeleteDemo = func(c *gin.Context) {
	c.String(200, "Delete demo router")
}

var JsonDemo = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // 💯
		"data": "<b>Hello world</b>",
	})
}

var PureJsonDemo = func(c *gin.Context) {
	c.PureJSON(http.StatusOK, gin.H{
		"data": "<b>Hello world</b>",
	})
}

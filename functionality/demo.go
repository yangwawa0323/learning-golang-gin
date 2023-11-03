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

// /frontend
// /frontend-v2
var FrontEnd = func(c *gin.Context) {
	version := c.DefaultQuery("version", "v1")
	if version == "v2" {
		c.Redirect(http.StatusPermanentRedirect, "/frontend-v2")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/frontend-v1")
}

var FrontEndV1 = func(c *gin.Context) {
	c.HTML(http.StatusOK, "v1.html", nil)
}
var FrontEndV2 = func(c *gin.Context) {
	c.HTML(http.StatusOK, "v2.html", nil)
}

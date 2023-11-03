package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func printRequest(req *http.Request) {
	log.Printf("Method: %#v\n", req.Method)
	log.Printf("URL path: %#v\n", req.URL.Path)

	log.Printf("Host: %#v\n", req.Host)
	log.Printf("Client address: %#v\n", req.RemoteAddr)
	log.Printf("Client request URI: %#v\n", req.RequestURI)
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		printRequest(c.Request)
		// END of my request in
		c.Next()
		// response out
		latency := time.Since(t)
		log.Print("latency: ", latency)

		status := c.Writer.Status()
		log.Print("other middleware return code :", status)
	}
}

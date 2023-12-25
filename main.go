package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Proxy endpoint 1
	endpoint1, _ := url.Parse("http://icanhazip.com")
	proxy1 := httputil.NewSingleHostReverseProxy(endpoint1)

	r.Use(middlewareDispatcher(proxy1))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}

func middlewareDispatcher(pr *httputil.ReverseProxy) gin.HandlerFunc {
	//get host

	return func(c *gin.Context) {
		host := c.Request.Host
		if host == "icanhazip.com" {
			pr.ServeHTTP(c.Writer, c.Request)
			return
		} else if host == "www.example2.com" {
			return
		} else {
			c.Status(http.StatusNotFound)
		}
	}
}

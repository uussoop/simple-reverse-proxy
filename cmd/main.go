package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/simple-reverse-proxy/configreader"
)

type proxyFromInStructure map[string]*httputil.ReverseProxy

func Initrouter(config *configreader.Config) {
	r := gin.Default()
	proxies := proxyFromInStructure{}
	for _, v := range config.Proxies {
		endpoint, _ := url.Parse(v.To.Scheme + v.To.Host + v.To.Port)
		proxy := httputil.NewSingleHostReverseProxy(endpoint)
		proxies[v.From.Host] = proxy
	}
	r.Use(middlewareDispatcher(&proxies))

	r.Any("/*any", func(c *gin.Context) {
		c.JSON(200, "nothing to see here")
	})

	r.Run()
}

func middlewareDispatcher(pr *proxyFromInStructure) gin.HandlerFunc {
	fmt.Println("middlewareDispatcher")
	return func(c *gin.Context) {
		host := c.Request.Host
		for k, v := range *pr {

			if strings.Contains(k, host) {

				v.ServeHTTP(c.Writer, c.Request)
				return
			}
		}

	}
}

func main() {
	config, err := configreader.ReadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	Initrouter(config)
}

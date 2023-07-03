package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func tryGetUserId(c *gin.Context) (string, bool) {
	authToken := c.GetHeader("X-Auth-Token")
	switch authToken {
	case "marcin":
		return "1", true
	default:
		return "", false
	}
}

func proxyHandler(remoteUrl *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := tryGetUserId(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		enrichedHeaders := c.Request.Header
		enrichedHeaders.Add("X-User-Id", userID)
		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		proxy.Director = func(req *http.Request) {
			req.Header = enrichedHeaders
			req.Host = remoteUrl.Host
			req.URL.Scheme = remoteUrl.Scheme
			req.URL.Host = remoteUrl.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	remoteUrl := os.Getenv("REMOTE_URL")
	url, err := url.Parse(remoteUrl)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.Any("/*proxyPath", proxyHandler(url))
	router.Run(":8080")
}

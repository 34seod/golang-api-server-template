package server

import (
	"fmt"
	"golang-api-server-template/configs"
	"golang-api-server-template/internal/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServerStart() {
	r := SetRouter()
	cnf := configs.Get()
	host := fmt.Sprintf("%s:%s", cnf.ServerHost, cnf.ServerPort)

	setSecureHeader(r, host)
	r.Run(fmt.Sprintf(":%s", cnf.ServerPort))
}

func SetRouter() *gin.Engine {
	r := gin.Default()
	routes.SetRoutes(r)
	return r
}

func setSecureHeader(r *gin.Engine, expectedHost string) {
	r.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost && c.Request.URL.Path != "/swagger/index.html" && !strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})
}

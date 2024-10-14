package routes

import (
	"net/http"
	"person/config"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env config.Env, timeout time.Duration, r *gin.Engine) {
	personRouter := r.Group("/person")
	NewPersonRoute(env, personRouter, timeout)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
	})
}

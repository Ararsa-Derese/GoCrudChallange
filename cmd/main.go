package main

import (
	"person/config"
	"person/delivery/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.NewEnv()
	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     env.CORSAllowMethods,
		AllowHeaders:     env.CORSAllowHeaders,
		ExposeHeaders:    env.CORSExposeHeaders,
		AllowCredentials: env.CORSAllowCredentials,
		MaxAge:           time.Duration(env.CORSMaxAge) * time.Second,
	}
	gin.Use(cors.New(corsConfig))
	routes.Setup(*env, timeout, gin)
	gin.Run()
}

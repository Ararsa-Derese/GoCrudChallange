package config

import (
	"os"
	"strconv"
	"strings"
)

type Env struct {
	AppEnv               string `mapstructure:"APP_ENV"`
	AppPort              string `mapstructure:"APP_PORT"`
	ContextTimeout       int    `mapstructure:"CONTEXT_TIMEOUT"`
	CORSAllowOrigins     []string
	CORSAllowMethods     []string
	CORSAllowHeaders     []string
	CORSExposeHeaders    []string
	CORSAllowCredentials bool
	CORSMaxAge           int
}

func NewEnv() *Env {
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	corsAllowCredentials, _ := strconv.ParseBool(os.Getenv("CORS_ALLOW_CREDENTIALS"))
	corsMaxAge, _ := strconv.Atoi(os.Getenv("CORS_MAX_AGE"))

	return &Env{
		AppEnv:               os.Getenv("APP_ENV"),
		AppPort:              os.Getenv("APP_PORT"),
		ContextTimeout:       contextTimeout,
		CORSAllowOrigins:     strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ","),
		CORSAllowMethods:     strings.Split(os.Getenv("CORS_ALLOW_METHODS"), ","),
		CORSAllowHeaders:     strings.Split(os.Getenv("CORS_ALLOW_HEADERS"), ","),
		CORSExposeHeaders:    strings.Split(os.Getenv("CORS_EXPOSE_HEADERS"), ","),
		CORSAllowCredentials: corsAllowCredentials,
		CORSMaxAge:           corsMaxAge,
	}
}

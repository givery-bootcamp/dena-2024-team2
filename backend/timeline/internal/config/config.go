package config

import (
	"os"
	"strconv"
)

var HostName = "127.0.0.1"
var Port = 9000
var CorsAllowOrigin = "http://localhost:3000"
var DBHostName = "db"
var DBUserName = "root"
var DBPassword = ""
var DBPort = 3306
var DBName = "training"
var JWTSecret = ""

func init() {
	if v := os.Getenv("HOSTNAME"); v != "" {
		HostName = v
	}
	if v, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64); err == nil {
		Port = int(v)
	}
	if v := os.Getenv("CORS_ALLOW_ORIGIN"); v != "" {
		CorsAllowOrigin = v
	}
	if v := os.Getenv("DB_HOSTNAME"); v != "" {
		DBHostName = v
	}
	if v := os.Getenv("DB_USERNAME"); v != "" {
		DBUserName = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		DBPassword = v
	}
	if v, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64); err == nil {
		DBPort = int(v)
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DBName = v
	}
	if v := os.Getenv("JWT_SECRET"); v != "" {
		JWTSecret = v
	}
}

package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CorsConfigDefault = cors.Config{
	Next:             nil,
	AllowOrigins:     "*",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           0,
}

// var CorsConfigDefault = cors.Config{
// 	Next:             nil,
// 	AllowOrigins:     "http://localhost:5173",
// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
// 	AllowHeaders:     "",
// 	AllowCredentials: true,
// 	ExposeHeaders:    "",
// 	MaxAge:           0,
// }

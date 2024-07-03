package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/router"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false, //Deploy ปรับเป็น true
	})
	app.Use(cors.New(config.CorsConfigDefault))

	database.ConnectDB()
	router.SetupRoutes(app)
	// app.Listen(":7000")
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func init() {
	initConfig()
	// initConfigEnv()
	// log.LogInit()
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// viper.AddConfigPath("/app")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }

}

// func initConfigEnv() {
// 	client := infisical.NewInfisicalClient(infisical.Config{
// 		SiteUrl: "http://10.255.253.90:8121", // ใช้ HTTP หรือ HTTPS ตามที่เซิร์ฟเวอร์รองรับ
// 	})

// 	// ประกาศตัวแปร err ก่อนใช้งาน
// 	var err error

// 	_, err = client.Auth().UniversalAuthLogin("4e58868a-3e79-4112-bddd-b3f9a16973e8", "d4f0957911541d2affb9e6601c34484e28531c2c75bd2ebe8f79d0d8b258dcba")

// 	if err != nil {
// 		fmt.Printf("Authentication failed: %v", err)
// 		os.Exit(1)
// 	}

// 	apiKeySecret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
// 		SecretKey:   "d4f0957911541d2affb9e6601c34484e28531c2c75bd2ebe8f79d0d8b258dcba",
// 		Environment: "backend",
// 		ProjectID:   "f6f0c47e-fe42-428e-ba44-d87ada01238f",
// 		SecretPath:  "/backend",
// 	})

// 	if err != nil {
// 		fmt.Printf("Error: %v", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("API Key Secret: %v", apiKeySecret)
// }

// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/puvadon-artmit/gofiber-template/config"
// 	"github.com/puvadon-artmit/gofiber-template/database"

// 	// "github.com/puvadon-artmit/gofiber-template/log"
// 	"github.com/puvadon-artmit/gofiber-template/router"
// 	"github.com/spf13/viper"
// )

// func main() {
// 	initConfig()

// 	app := fiber.New(fiber.Config{
// 		Prefork: false, //Deploy ปรับเป็น true
// 	})
// 	app.Use(cors.New(config.CorsConfigDefault))

// 	database.ConnectDB()
// 	// log.LogInit(app) // Pass the app to LogInit
// 	router.SetupRoutes(app)

// 	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
// }

// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
// 	viper.AutomaticEnv()
// 	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

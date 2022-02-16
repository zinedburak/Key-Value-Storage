package main

import (
	"Key_Value_Storage/controllers"
	"Key_Value_Storage/models"
	"Key_Value_Storage/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

var Db models.Store

func main() {
	Db := models.CreateStore("HelloWorld.json", 5*time.Second)
	controllers.Db = Db

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	routes.Setup(app)

	app.Listen(":8000")
}

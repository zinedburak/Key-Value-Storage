package routes

import (
	"Key_Value_Storage/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/get-key-value", controllers.GetKeyValue)
	app.Post("/api/set-key-value", controllers.SetKeyValue)
	app.Get("/api/get-all-key-value", controllers.GetAllKeyValue)
	app.Get("/api/save-all-key-value", controllers.FlushAllData)
}

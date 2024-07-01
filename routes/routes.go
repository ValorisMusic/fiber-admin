package routes

import (
	"fiber-admin/controllers"

	"github.com/gofiber/fiber/v3"
)

func UserRouts(app *fiber.App) {
	User := app.Group("/api")
	{

		User.Post("/GetLevelByID", controllers.GetLevelByID)
		User.Post("/GetLevel", controllers.GetLevel)
		User.Post("/CreateLevel", controllers.CreateLevel)
		User.Post("/DeleteLevelByID", controllers.DeleteLevelByID)
		User.Post("/UpdateLevelByID", controllers.UpdateLevelByID)

	}
}

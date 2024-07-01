package main

import (
	"fiber-admin/initializers"
	"fiber-admin/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func init() {
	initializers.LoadEnvVariables()
	initializers.DBconnect()
	initializers.InitRedis()

}
func main() {

	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	routes.UserRouts(app)

	app.Listen(":3000")

}

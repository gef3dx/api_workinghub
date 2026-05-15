// Сборка приложения
package bootstrap

import "github.com/gofiber/fiber/v3"

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		//Настройки Fiber
		AppName: "api_workinghub",
	})
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	return app
}

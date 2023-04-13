package swagger

import (
	_ "golang/external/fiber/swagger/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func App(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}

// TODO: Markdown + Example Response + Example Code

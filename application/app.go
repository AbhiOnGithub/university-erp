package application

import (
	"fmt"
	"github.com/abhiongithub/university-erp/database"
	"github.com/abhiongithub/university-erp/routers"
	"github.com/gofiber/fiber"
)

type App struct {
	apiVersion string
}

func SetupApp() {
	app := App{  }
	app.apiVersion = "v1" //os.Getenv("API_Version")

	database.InitDatabase()

	fiberApp := fiber.New()

	courseRouter := routers.NewCourseRouter()
	courseRouter.SetupCourseRoutes(fiberApp, app.apiVersion)

	fiberApp.Listen(3000)

	fmt.Println("Setting up application")
}

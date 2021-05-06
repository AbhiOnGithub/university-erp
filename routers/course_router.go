package routers

import (
	"github.com/abhiongithub/university-erp/repositories"
	"github.com/gofiber/fiber"
)

type CourseRouter struct{

}

func NewCourseRouter() *CourseRouter {
	return &CourseRouter{}
}

func (courseRouter *CourseRouter) SetupCourseRoutes(app *fiber.App, version string) {
	courseRepo := repositories.NewCourseRepository()

	app.Get("/api/"+version+"/course", courseRepo.GetAllCourses)
	app.Post("/api/"+version+"/course", courseRepo.AddCourse)
	app.Get("/api/"+version+"/course/:id", courseRepo.GetCourse)
	app.Delete("/api/"+version+"/course/:id", courseRepo.DeleteCourse)

}
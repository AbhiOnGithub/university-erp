package repositories

import (
	"github.com/abhiongithub/university-erp/database"
	"github.com/abhiongithub/university-erp/models"
	"github.com/gofiber/fiber"
)

type CourseRepository struct {
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

//GetAllCourses
func (courseRepo *CourseRepository) GetAllCourses(ctx *fiber.Ctx){
	var db = database.DBConn
	var courses = []models.Course{}
	db.Find(&courses)
	ctx.JSON(courses)
}

//GetCourse
func (courseRepo *CourseRepository) GetCourse(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var course models.Course
	db.Find(&course, id)
	c.JSON(course)
}

//AddCourse
func (courseRepo *CourseRepository) AddCourse(c *fiber.Ctx) {
	db := database.DBConn
	course := new(models.Course)
	if err := c.BodyParser(course); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&course)
	c.JSON(course)
}

func (courseRepo *CourseRepository)  DeleteCourse(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var course models.Course
	db.First(&course, id)
	if course.Name == "" {
		c.Status(500).Send("No Course Found with ID")
		return
	}
	db.Delete(&course)
	c.Send("Course Successfully deleted")
}


/*{
	models.Course{Id: "13", Name: "MS", Details: "Masters in Computer Science", AdmissionsOpen: false, NoOfSemesters: 4, PreviousBatches: nil},
	models.Course{Id: "14", Name: "BS", Details: "Bachelors in Cyber Security", AdmissionsOpen: true, NoOfSemesters: 6, PreviousBatches: nil},
}*/
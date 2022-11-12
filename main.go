package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task_2/controller"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	studentController := controller.StudentController{}
	e.GET("/students", studentController.GetAllStudents)
	e.POST("/create_students", studentController.Create)
	e.GET("/students/:id", studentController.GetStudent)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.PATCH("/students/:id", studentController.PatchStudent)
	e.DELETE("/students/:id", studentController.DeleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

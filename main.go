package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task_2/controller"
	"task_2/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	dbConnection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=helloyzsf dbname=students port=5432"), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")

	// Routes
	studentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: dbConnection,
		},
	}
	e.GET("/students", studentController.GetAllStudents)
	e.POST("/create_students", studentController.Create)
	e.GET("/students/:id", studentController.GetStudent)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.PATCH("/students/:id", studentController.PatchStudent)
	e.DELETE("/students/:id", studentController.DeleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

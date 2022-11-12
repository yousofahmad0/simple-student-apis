package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
	"task_2/entity"
	"task_2/repository"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

var (
	lock = sync.Mutex{}
)

//----------
// Handlers
//----------

func (r StudentController) GetAllStudents(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, r.StudentRepository.GetAll())
}

func (r StudentController) GetStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entity.Student)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	res, s := r.StudentRepository.GetStudent(id)
	if s == nil {
		return c.JSON(http.StatusNotFound, res)
	}
	return c.JSON(http.StatusOK, s)
}

func (r StudentController) Create(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &entity.Student{}
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, r.StudentRepository.Create(u))
}

func (r StudentController) UpdateStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(entity.Student)
	if err := c.Bind(u); err != nil {
		return err
	}
	res, s := r.StudentRepository.Update(id, u)
	if s == nil {
		return c.JSON(http.StatusNotFound, res)
	}
	return c.JSON(http.StatusOK, s)
}

func (r StudentController) PatchStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(entity.Student)
	if err := c.Bind(u); err != nil {
		return err
	}
	res, s := r.StudentRepository.Patch(id, u)
	if s == nil {
		return c.JSON(http.StatusNotFound, res)
	}
	return c.JSON(http.StatusOK, s)
}

func (r StudentController) DeleteStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	if r.StudentRepository.Delete(id) == "deleted successfully" {
		return c.JSON(http.StatusAccepted, "deleted successfully")
	}
	return c.JSON(http.StatusNotFound, "student not found")
}

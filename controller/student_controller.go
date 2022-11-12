package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
	"task_2/entity"
)

type StudentController struct {
}

var (
	students = map[int]*entity.Student{}
	seq      = 1
	lock     = sync.Mutex{}
)

//----------
// Handlers
//----------

func (receiver StudentController) GetAllStudents(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, students)
}

func (receiver StudentController) GetStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, students[id])
}

func (receiver StudentController) Create(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &entity.Student{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	students[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func (receiver StudentController) UpdateStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entity.Student)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	students[id].FirstName = u.FirstName
	students[id].LastName = u.LastName
	return c.JSON(http.StatusOK, students[id])
}

func (receiver StudentController) PatchStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(entity.Student)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if u.FirstName != "" {
		students[id].FirstName = u.FirstName
	}
	if u.LastName != "" {
		students[id].LastName = u.LastName
	}
	return c.JSON(http.StatusOK, students[id])
}

func (receiver StudentController) DeleteStudent(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(students, id)
	return c.NoContent(http.StatusNoContent)
}

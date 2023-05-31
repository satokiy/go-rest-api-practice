package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// interface definition
type ITaskController interface {
	CreateTask(c echo.Context) error
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

// constructor
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{
		tu: tu,
	}
}

// CreateTask implements ITaskController
func (tc *taskController) CreateTask(c echo.Context) error {
	// contextからuserを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))

	res, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}

// DeleteTask implements ITaskController
func (tc *taskController) DeleteTask(c echo.Context) error {
	panic("unimplemented")
}

// GetAllTasks implements ITaskController
func (tc *taskController) GetAllTasks(c echo.Context) error {
	panic("unimplemented")
}

// GetTaskById implements ITaskController
func (tc *taskController) GetTaskById(c echo.Context) error {
	panic("unimplemented")
}

// UpdateTask implements ITaskController
func (tc *taskController) UpdateTask(c echo.Context) error {
	panic("unimplemented")
}

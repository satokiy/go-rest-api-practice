package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

// routerののコンストラクタ。ここではcontrollerをDIする
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	
	// user
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
    // task
	e.POST("/tasks", tc.CreateTask)

	return e
}

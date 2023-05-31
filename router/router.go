package router

import (
	"fmt"
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
  }
  
// routerののコンストラクタ。ここではcontrollerをDIする
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.BodyDump(bodyDumpHandler))
	
	// user
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
    // task
	e.POST("/tasks", tc.CreateTask)

	return e
}

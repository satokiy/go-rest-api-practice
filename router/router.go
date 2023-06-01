package router

import (
	"fmt"
	"go-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
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
	// add middleware
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// user
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)

	// task group
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	e.GET("", tc.GetAllTasks)
	e.GET("/:taskId", tc.GetTaskById)
	e.POST("", tc.CreateTask)
	e.PUT("/:taskId", tc.UpdateTask)
	e.DELETE("/:taskId", tc.DeleteTask)

	return e
}

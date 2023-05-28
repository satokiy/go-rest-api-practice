package controller

// 本pjtではechoを使う
import (
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
)

// controllerの定義
// routerから呼び出され、usecaseを呼び出す
type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// controllerのコンストラクタ
// controllerに対してusecaseをDIする
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{
		uu: uu,
	}
}

// Login implements IUserController
func (uc *userController) Login(c echo.Context) error {
	// TODO: implement
	panic("unimplemented")
}

// Logout implements IUserController
func (uc *userController) Logout(c echo.Context) error {
	// TODO: implement
	panic("unimplemented")
}

// SignUp implements IUserController
func (uc *userController) SignUp(c echo.Context) error {
	// TODO: implement
	panic("unimplemented")
}

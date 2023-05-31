package controller

// 本pjtではechoを使う
import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

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
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// cookieをセット
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Domain = os.Getenv("API_DOMAIN")
	// 検証用にコメントアウト
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode // cross site
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// Logout implements IUserController
func (uc *userController) Logout(c echo.Context) error {
	// cookieを削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.Expires = time.Now()
	cookie.Domain = os.Getenv("API_DOMAIN")
	// 検証用にコメントアウト
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode // cross site
	c.SetCookie(cookie)
	
	return c.NoContent(http.StatusOK)
}

// SignUp implements IUserController
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

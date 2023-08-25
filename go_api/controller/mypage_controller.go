package controller

import (
	"go_api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IMypageController interface {
	GetUser(c echo.Context) error
}

type mypageController struct {
	mu usecase.IMypageUsecase
}

func NewMypageContorller(mu usecase.IMypageUsecase) IMypageController {
	return &mypageController{mu}
}

func (mc *mypageController) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// jwtの情報から、ログインしているuserのidを取得。
	userId := claims["user_id"]

	userRes, err := mc.mu.GetUser(uint(userId.(float64)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)

}

package controller

import (
	"net/http"

	"github.com/barbaromatrix/clean-example/domain/model"
	"github.com/barbaromatrix/clean-example/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.GET(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

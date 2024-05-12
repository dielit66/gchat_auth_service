package controller

import (
	"net/http"

	"github.com/xxlifestyle/auth_service/internal/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func (uc *userController) GetUser(w http.ResponseWriter, r *http.Request) {

}
func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func NewUserController(userInteractor interactor.UserInteractor) UserController {
	return &userController{
		userInteractor: userInteractor,
	}
}

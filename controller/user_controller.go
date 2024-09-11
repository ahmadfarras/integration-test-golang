package controller

import (
	"encoding/json"
	"farras/integration-test-golang/model"
	"farras/integration-test-golang/usecase"
	"net/http"
)

type UserController struct {
	userUsecase usecase.UserUseCase
}

func NewUserController(userUsecase usecase.UserUseCase) *UserController {
	return &UserController{userUsecase}
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := c.userUsecase.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCreated := c.userUsecase.Create(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userCreated)
}

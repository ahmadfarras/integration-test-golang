package main

import (
	"farras/integration-test-golang/controller"
	"farras/integration-test-golang/model"
	"farras/integration-test-golang/repository"
	"farras/integration-test-golang/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	var db *gorm.DB
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1:%s)/belajar_golang?charset=utf8mb4&parseTime=True&loc=Local", port)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}

	db.AutoMigrate(&model.User{})

	r := chi.NewRouter()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	r.Get("/users", userController.GetUsers)
	r.Post("/users", userController.CreateUser)

	http.ListenAndServe(":8080", r)
}

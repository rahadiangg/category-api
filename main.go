package main

import (
	"log"
	"net/http"
	"os"
	"rahadiangg/category-api/app"
	"rahadiangg/category-api/controller"
	"rahadiangg/category-api/helper"
	"rahadiangg/category-api/repository"
	"rahadiangg/category-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("APP_PORT"),
		Handler: router,
		// Handler: middleware.NewAuthMiddleware(router),
	}
	log.Println("App run on port ", os.Getenv("APP_PORT"))

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

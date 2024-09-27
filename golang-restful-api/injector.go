//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10" // validator
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func ProvideValidator() *validator.Validate {
	v := validator.New()
	// Contoh: menambahkan custom validator atau pengaturan lainnya
	return v
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		ProvideValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

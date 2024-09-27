package main

import (
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"

	_ "github.com/lib/pq"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

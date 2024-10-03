package server

import (
	"net/http"
	"trivia-app/api/middleware"
	"trivia-app/api/router"
)

func Server() *http.Server {
	router := router.Router()
	middlewareStack := middleware.Stack()

	server := http.Server{
		Addr:    ":8080",
		Handler: middlewareStack(router),
	}

	return &server
}

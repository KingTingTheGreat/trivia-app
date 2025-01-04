package server

import (
	"net/http"
	"trivia-app/middleware"
	"trivia-app/router"
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

package controllers

import (
	"gorm_example/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/user", UserController)
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {})
	return http.ListenAndServe(":"+config.Config.ServerPort, nil)
}

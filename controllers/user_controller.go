package controllers

import (
	"encoding/json"
	"fmt"
	"gorm_example/models"
	"log"
	"net/http"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		var users []models.User
		models.GetAllUsers(&users)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		fmt.Println("POST")
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
		}
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		models.InsertUser(&user)
		w.WriteHeader(http.StatusCreated)
		// 最後にしないとstatus code が反映されない
		json.NewEncoder(w).Encode(user)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/riffypatel/todo-api/models"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User
var jwtKey = []byte("secret_key")

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)
	user.ID = len(users) + 1
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})

}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	var storedUser *models.User
	for _, u := range users {
		if u.Username == creds.Username {
			storedUser = &u
			break
		}
	}
	if storedUser == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": storedUser.Username,
		"user_id":  storedUser.ID,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

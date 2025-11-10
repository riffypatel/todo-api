package main
import (
	"log"
	"net/http"
	"os"

	"todo-api/handlers"
	"todo-api/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register",handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	api := router.PathPrefix("/todos").Subrouter()
	api.Use(middleware.JWTAuth)
	api.HandleFunc("", handlers.CreateTodo).Methods("POST")
	api.HandleFunc("", handlers.GetTodos).Methods("GET")
	api.HandleFunc("/{id}", handlers.UpdateTodo).Methods("PUT")
	api.HandleFunc("/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server running on port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

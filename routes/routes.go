package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"postgres/services"
)

func CreateRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", services.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts", services.CreatePost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id}", services.UpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/posts/{id}", services.DeletePost).Methods(http.MethodDelete)

	return router
}

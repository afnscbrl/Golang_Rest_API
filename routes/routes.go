package routes

import (
	"log"
	"net/http"

	"github.com/afnscbrl/Golang_Rest_API/controller"
	"github.com/afnscbrl/Golang_Rest_API/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/dashboard", controller.Dashboard).Methods("Get")
	r.HandleFunc("/api/receitas", controller.NewIncome).Methods("Post")
	r.HandleFunc("/api/receitas", controller.Income).Methods("Get")
	r.HandleFunc("/api/receitas/", controller.Income).Methods("Get")
	r.HandleFunc("/api/receitas/{id}", controller.IncomeDetail).Methods("Get")
	r.HandleFunc("/api/receitas/{id}", controller.EditIncome).Methods("Put")
	r.HandleFunc("/api/receitas/{id}", controller.DeleteIncome).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

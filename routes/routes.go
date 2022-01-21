package routes

import (
	"log"
	"net/http"

	"github.com/afnscbrl/Golang_Rest_API/controller"
	"github.com/afnscbrl/Golang_Rest_API/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//This function contains all URL routes of API and the html methods.
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
	r.HandleFunc("/api/despesas", controller.NewOutcome).Methods("Post")
	r.HandleFunc("/api/despesas", controller.Outcome).Methods("Get")
	r.HandleFunc("/api/despesas/", controller.Outcome).Methods("Get")
	r.HandleFunc("/api/despesas/{id}", controller.OutcomeDetail).Methods("Get")
	r.HandleFunc("/api/despesas/{id}", controller.EditOutcome).Methods("Put")
	r.HandleFunc("/api/despesas/{id}", controller.DeleteOutcome).Methods("Delete")
	//set port 8000 to list and serve
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

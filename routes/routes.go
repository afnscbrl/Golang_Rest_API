package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/afnscbrl/Golang_Rest_API/controller"
	"github.com/afnscbrl/Golang_Rest_API/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//This function contains all URL routes of API and the html methods.
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home).Methods("Get")
	r.HandleFunc("/login", controller.Login).Methods("Post")
	r.HandleFunc("/register", controller.Register).Methods("Post")
	//Quando implementar o front, editar o Middleware

	subRouter := r.PathPrefix("/api/").Subrouter()
	subRouter.Use(middleware.ContentTypeMiddleware, middleware.AuthorizationMiddleware)
	// r.HandleFunc("/api/dashboard", controller.Dashboard).Methods("Get")
	subRouter.HandleFunc("/resumo/{year}/{month}", controller.BalanceByMonth).Methods("Get")
	subRouter.HandleFunc("/receitas", controller.NewIncome).Methods("Post")
	subRouter.HandleFunc("/receitas", controller.Income).Methods("Get")
	subRouter.HandleFunc("/receitas/{id}", controller.IncomeDetail).Methods("Get")
	subRouter.HandleFunc("/receitas/{id}", controller.EditIncome).Methods("Put")
	subRouter.HandleFunc("/receitas/{id}", controller.DeleteIncome).Methods("Delete")
	subRouter.HandleFunc("/receitas/{year}/{month}", controller.IncomeByMonth).Methods("Get")
	subRouter.HandleFunc("/despesas", controller.NewOutcome).Methods("Post")
	subRouter.HandleFunc("/despesas", controller.Outcome).Methods("Get")
	subRouter.HandleFunc("/despesas/{id}", controller.OutcomeDetail).Methods("Get")
	subRouter.HandleFunc("/despesas/{id}", controller.EditOutcome).Methods("Put")
	subRouter.HandleFunc("/despesas/{id}", controller.DeleteOutcome).Methods("Delete")
	subRouter.HandleFunc("/despesas/{year}/{month}", controller.OutcomeByMonth).Methods("Get")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	//set port 8000 to list and serve
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

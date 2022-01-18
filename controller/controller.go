package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/routes/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	var res []models.Resume
	// var inc []models.Income
	// var outp []models.OutcomePer
	// var outv []models.OutcomeVar
	// var bal []models.Balance
	database.DB.Find(&res)
	json.NewEncoder(w).Encode(res)
}

func NewIncome(w http.ResponseWriter, r *http.Request) {
	var newIncome models.Income
	json.NewDecoder(r.Body).Decode(&newIncome)
	database.DB.Create(&newIncome)
	json.NewEncoder(w).Encode(newIncome)
}

func Income(w http.ResponseWriter, r *http.Request) {
	var inc []models.Income
	database.DB.Find(&inc)
	json.NewEncoder(w).Encode(inc)
}

func IncomeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var income models.Income
	database.DB.First(&income, id)
	json.NewEncoder(w).Encode(income)
}

func EditIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var income models.Income
	database.DB.First(&income, id)
	json.NewDecoder(r.Body).Decode(&income)
	database.DB.Save(&income)
	json.NewEncoder(w).Encode(income)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var income models.Income
	database.DB.Delete(&income, id)
	json.NewEncoder(w).Encode(income)
}

func NewOutcome(w http.ResponseWriter, r *http.Request) {
	var newOutcome models.Outcome
	json.NewDecoder(r.Body).Decode(&newOutcome)
	database.DB.Create(&newOutcome)
	json.NewEncoder(w).Encode(newOutcome)
}

func Outcome(w http.ResponseWriter, r *http.Request) {
	var out []models.Outcome
	database.DB.Find(&out)
	json.NewEncoder(w).Encode(out)
}

func OutcomeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var outcome models.Outcome
	database.DB.First(&outcome, id)
	json.NewEncoder(w).Encode(outcome)

}

func EditOutcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var outcome models.Outcome
	database.DB.First(&outcome, id)
	json.NewDecoder(r.Body).Decode(&outcome)
	database.DB.Save(&outcome)
	json.NewEncoder(w).Encode(outcome)
}

func DeleteOutcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var outcome models.Outcome
	database.DB.Delete(&outcome, id)
	json.NewEncoder(w).Encode(outcome)
}

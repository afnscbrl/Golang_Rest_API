package controller

import (
	"encoding/json"
	"net/http"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/models"
	"github.com/gorilla/mux"
)

func BalanceByMonth(w http.ResponseWriter, r *http.Request) {
	var (
		total_income  float64
		total_outcome float64
		outcome       []models.Outcome
		income        []models.Income
	)

	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]

	database.DB.Where("year = ? AND month = ?", year, month).Find(&outcome)
	database.DB.Where("year = ? AND month = ?", year, month).Find(&income)

	for i := 0; i < len(income); i++ {
		total_income += income[i].Value
	}

	for j := 0; j < len(outcome); j++ {
		total_outcome += outcome[j].Value
	}

	res := models.Resume{
		TotalIncome:  total_income,
		TotalOutcome: total_outcome,
		Balance:      (total_income - total_outcome),
		//add by category
	}

	json.NewEncoder(w).Encode(res)
}

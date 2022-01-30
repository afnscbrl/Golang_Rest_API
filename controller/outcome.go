package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/models"
	"github.com/gorilla/mux"
)

//This controller contains all func that control the in/out data from requests and db

// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Home Page")
// }

// func Dashboard(w http.ResponseWriter, r *http.Request) {
// 	var res []models.Resume
// 	// var inc []models.Income
// 	// var outp []models.OutcomePer
// 	// var outv []models.OutcomeVar
// 	// var bal []models.Balance
// 	database.DB.Find(&res)
// 	json.NewEncoder(w).Encode(res)
// }

func NewOutcome(w http.ResponseWriter, r *http.Request) {
	var newOutcome models.Outcome
	var outcome []models.Outcome

	json.NewDecoder(r.Body).Decode(&newOutcome)

	if newOutcome.Id != 0 {
		log.Println("Error 400: Bad Request - Id don't be set")
		http.Error(w, badResquest+" - Id don't be set", http.StatusBadRequest)
		return
	}

	if newOutcome.Category == "" {
		newOutcome.Category = "Outras"
	}
	dates := strings.Split(newOutcome.Date, "-")
	day, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	year, _ := strconv.Atoi(dates[0])
	if (day < 1 || day > 31) || (month < 1 || month > 12) {
		log.Println("Error 400: BAD RESQUEST - Wrong data type")
		http.Error(w, badResquest+" - Wrong data type", http.StatusBadRequest)
		return
	}
	newOutcome.Year = year
	newOutcome.Month = month
	newOutcome.Day = day

	database.DB.Where("month = ?", month).Find(&outcome)
	for i := 0; i < len(outcome); i++ {
		if outcome[i].Describe == newOutcome.Describe {
			log.Println("Error 409: Conflict - This desccribe already exist in this month")
			http.Error(w, conflict+" - This desccribe already exist for this month", http.StatusConflict)
			return
		}
	}
	database.DB.Create(&newOutcome)
	json.NewEncoder(w).Encode(newOutcome)
}

func Outcome(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()
	if !v.Has("describe") {
		var out []models.Outcome
		database.DB.Find(&out)
		json.NewEncoder(w).Encode(out)

	} else {
		desc := v.Get("describe")
		var outcome models.Outcome
		database.DB.Where("describe = ?", desc).First(&outcome)
		json.NewEncoder(w).Encode(outcome)
	}
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
	var outcomeDB []models.Outcome
	database.DB.First(&outcome, id)
	json.NewDecoder(r.Body).Decode(&outcome)
	dates := strings.Split(outcome.Date, "-")
	day, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	year, _ := strconv.Atoi(dates[0])
	if (day < 1 || day > 31) || (month < 1 || month > 12) {
		log.Println("Error 400: BAD RESQUEST - Wrong data type")
		http.Error(w, badResquest+" - Wrong data type", http.StatusBadRequest)
		return
	}
	outcome.Year = year
	outcome.Month = month
	outcome.Day = day

	database.DB.Where("month = ?", month).Find(&outcomeDB)
	for i := 0; i < len(outcomeDB); i++ {
		if outcomeDB[i].Describe == outcome.Describe {
			log.Println("Error 409: Conflict - This desccribe already exist in this month")
			http.Error(w, conflict+" - This desccribe already exist for this month", http.StatusConflict)
			return
		}
	}
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

func OutcomeByMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]

	var outcome []models.Outcome
	database.DB.Where("year = ? AND month = ?", year, month).Find(&outcome)
	json.NewEncoder(w).Encode(outcome)
}

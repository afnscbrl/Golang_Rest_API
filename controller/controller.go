package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/models"
	"github.com/gorilla/mux"
)

//This controller contains all func that control the in/out data from requests and db

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
	// r.ParseForm()
	// var err error
	// var newIncome models.Income
	// // get the value in html tag description
	// newIncome.Describe = r.PostForm.Get("description")
	// newIncome.Value, err = strconv.ParseFloat(r.PostForm.Get("value"), 64)
	// if err != nil {
	// 	fmt.Println("Error ParseFloat")
	// }
	// newIncome.Date = r.PostForm.Get("date")
	// database.DB.Create(&newIncome)
	// json.NewEncoder(w).Encode(newIncome)

	var newIncome models.Income
	json.NewDecoder(r.Body).Decode(&newIncome)
	database.DB.Create(&newIncome)
	json.NewEncoder(w).Encode(newIncome)
}

func Income(w http.ResponseWriter, r *http.Request) {
	// var inc []models.Income
	// var tmpl = template.Must(template.ParseGlob("./view/income.html"))
	// tmpl.ExecuteTemplate(w, "Income", inc)

	v := r.URL.Query()

	if !v.Has("describe") {
		var inc []models.Income
		database.DB.Find(&inc)
		json.NewEncoder(w).Encode(inc)
	} else {
		desc := v.Get("describe")
		var income models.Income
		database.DB.Where("describe = ?", desc).First(&income)
		json.NewEncoder(w).Encode(income)
	}

	//for pra listar todos os []inc e jogar no tamplate
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
	if newOutcome.Category == "" {
		newOutcome.Category = "Outras"
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

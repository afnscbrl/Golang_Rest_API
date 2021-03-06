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

var badResquest string
var conflict string

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

	badResquest = "Error 400 - Bad Request"
	conflict = "Error 409 - Conflict"

	var newIncome models.Income
	var income []models.Income
	json.NewDecoder(r.Body).Decode(&newIncome)

	if newIncome.Id != 0 {
		log.Println("Error 400: Bad Request - Id don't be set")
		http.Error(w, badResquest+" - Id don't be set", http.StatusBadRequest)
		return
	}

	dates := strings.Split(newIncome.Date, "-")
	day, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	year, _ := strconv.Atoi(dates[0])
	if (day < 1 || day > 31) || (month < 1 || month > 12) {
		log.Println("Error 400: BAD RESQUEST - Wrong data type")
		http.Error(w, badResquest+" - Wrong data type", http.StatusBadRequest)
		return
	}
	newIncome.Year = year
	newIncome.Month = month
	newIncome.Day = day

	database.DB.Where("month = ?", month).Find(&income)
	for i := 0; i < len(income); i++ {
		if income[i].Describe == newIncome.Describe {
			log.Println("Error 409: Conflict - This desccribe already exist in this month")
			http.Error(w, conflict+" - This desccribe already exist for this month", http.StatusConflict)
			return
		}
	}
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
	if income.Id == 0 {
		http.Error(w, "Record not found ", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(income)
}

func EditIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var income models.Income
	var incomeDB []models.Income
	database.DB.First(&income, id)
	if income.Id == 0 {
		http.Error(w, "Record not found ", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&income)
	dates := strings.Split(income.Date, "-")
	day, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	year, _ := strconv.Atoi(dates[0])
	if (day < 1 || day > 31) || (month < 1 || month > 12) {
		log.Println("Error 400: BAD RESQUEST - Wrong data type")
		http.Error(w, badResquest+" - Wrong data type", http.StatusBadRequest)
		return
	}
	income.Year = year
	income.Month = month
	income.Day = day

	database.DB.Where("month = ?", month).Find(&incomeDB)
	for i := 0; i < len(incomeDB); i++ {
		if incomeDB[i].Describe == income.Describe {
			log.Println("Error 409: Conflict - This desccribe already exist in this month")
			http.Error(w, conflict+" - This desccribe already exist for this month", http.StatusConflict)
			return
		}
	}

	database.DB.Save(&income)
	json.NewEncoder(w).Encode(income)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var income models.Income
	if income.Id == 0 {
		http.Error(w, "Record not found ", http.StatusNotFound)
		return
	}
	database.DB.Delete(&income, id)
	json.NewEncoder(w).Encode(income)
}

func IncomeByMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	int_month, _ := strconv.Atoi(month)
	if int_month <= 0 || int_month > 12 {
		log.Println("Error 400: BAD RESQUEST - Wrong data type")
		http.Error(w, badResquest+" - Wrong data type", http.StatusBadRequest)
		return
	}
	var income []models.Income
	database.DB.Where("year = ? AND month = ?", year, month).Find(&income)
	if len(income) == 0 {
		log.Println("Error 404: NOT FOUND")
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(income)

}

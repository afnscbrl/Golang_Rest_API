Intro:
Development of a REST API with Golang from the Alura challenge that i participated in. #alurachallengeback2

Languages and techs:
Golang with gorilla mux and gorm, Postgresql and Docker.

Motivation and Goal:
It was my first project in backend area and i choose golang because i really enjoy conding with this language. The goal was create a personal finance control that the requests and response were done with json.

Phases:
The challenge was divided in three weeks that simulated a real enviroment of a business demand. In the all project, i used the MVC pattern.

The first week - the tasks was create the basis of database with id, describe, value and date, so i choose Postgresql. After it, i need to create a CRUD (create, read, update and delete) routes for incomes and outcomes with one rule: the user didn't POST the same describe in the same month. Then, i created an constrain in database to prevent the user to do that and also in the development the model part, i verified if the rules was being followed, if didn't, i returned an 409 error - Conflict passing the message "This desccribe already exist in this month", as can seen below:

	for i := 0; i < len(income); i++ {
		if income[i].Describe == newIncome.Describe {
			log.Println("Error 409: Conflict - This desccribe already exist in this month")
			http.Error(w, conflict+" - This desccribe already exist for this month", http.StatusConflict)
			return
		}
	}
  
  In that week, i needed to create the routes to list all incomes and outcomes or one of then by id. 
  
  The second week - in that week the goals was to put category for outcomes POST and if this data was sent in blank i needed set automatically the category as "Other", another goals was to create a route that showed the incomes or outcomes for a specifically month and finally to create a resume route that brings the total of income, outcome, and the value gained (income) or spent (outocome) by category. For that week, i need make changes in database, create a new table with the categorys to improve de incomes and outcomes. To show all that resume, i need crate other table like this 
  
  create table resumes (
    id serial primary key,
    total_income real,
    total_outcome real,
    balance real
);

  That week like the first one wasn't challenge to me cause i could achieve all the demands whitout much struggles. My knowledge in sql and golang was suficient to finish the work.
  
  The third week - 
  
  

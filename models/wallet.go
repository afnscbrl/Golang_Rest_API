package models

type Resume struct {
	Id           int     `json:"id"`
	TotalIncome  float64 `json:"total income"`
	TotalOutcome float64 `json:"total outcome"`
	Balance      float64 `json:"balance"`
}

var Resumes []Resume

type Income struct {
	Id       int     `json:"id"`
	Describe string  `json:"describe"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
}

var Incomes []Income

type Outcome struct {
	Id       int     `json:"id"`
	Describe string  `json:"describe"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
}

var Outcomes []Income

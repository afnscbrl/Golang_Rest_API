package models

type Resume struct {
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
	Year     int     `json:"-"`
	Month    int     `json:"-"`
	Day      int     `json:"-"`
}

var Incomes []Income

type Outcome struct {
	Id       int     `json:"id"`
	Describe string  `json:"describe"`
	Value    float64 `json:"value"`
	Date     string  `json:"date"`
	Category string  `json:"category"`
	Year     int     `json:"-"`
	Month    int     `json:"-"`
	Day      int     `json:"-"`
}

var Outcomes []Income

type Category struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

var Categorys []Category

type Users struct {
	Id           int    `json:"-"`
	Username     string `json:"username"`
	Passwordhash string `json:"password"`
	Isdisable    bool   `json:"-"`
}

var Userss []Users

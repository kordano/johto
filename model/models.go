package model

import (
	"time"
)

// EffortUnit provides enum base
type EffortUnit int

const (
	//EffortDay provides enum for effort in a day
	EffortDay EffortUnit = 1 + iota
	//EffortMonth provides enum for effort in a month
	EffortMonth
	//EffortYear provides enum for effort in a year
	EffortYear
)

// Member provides data model for sql member
type Member struct {
	ID        int    `json:"_id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Task provides data model for sql task
type Task struct {
	ID         int
	Assignee   *Member
	Title      string
	Effort     float32
	EffortUnit EffortUnit
	Category   string
}

// Customer provides data model for sql customer
type Customer struct {
	ID         int `json:"_id"`
	Name       string `json: "name"`
	Contact    string `json: "contact"`
	Department string `json: "department"`
	Street     string `json: "street"`
	City       string `json: "city"`
	Postal     string `json: "postal"`
	Country    string `json: "country"`
}

// Document provides data model for sql document
type Document struct {
	ID        int
	timestamp time.Time
	Reference string
	FileName  string
}

// Project provides data model for sql project
type Project struct {
	ID          int
	StartDate   time.Time
	EndDate     time.Time
	Title       string
	Customer    *Customer
	Members     []*Member
	Responsible *Member
	AcceptedAt  time.Time
	PaidAt      time.Time
	Offers      []*Document
	Invoice     *Document
}

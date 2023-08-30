package models

import (
	"time"
)

type Todo struct {
	Id int
	Title string
	Completed int
	Color string
	Start time.Time
	End time.Time
}
package types

import "time"

//Page data object for handling page configuration in view
type Page struct {
	Title    string
	Messages []string
}

//Entry data object to contain duck entry
type Entry struct {
	ID            string
	TimeFed       time.Time
	Food          Food
	Location      string
	AmountOfFood  int
	NumberOfDucks int
}

//Food data object to contain food meta
type Food struct {
	Kind string
	Name string
}

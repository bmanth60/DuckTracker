package types

import "time"

type Page struct {
	Title    string
	Token    string
	Messages []string
}

type Entry struct {
	ID            string
	TimeFed       time.Time
	Food          Food
	Location      string
	AmountOfFood  int
	NumberOfDucks int
}

type Food struct {
	Kind string
	Name string
}

package data

import (
	"database/sql"

	"github.com/bmanth60/DuckTracker/types"
)

//Set function declaration for adding a data set
type Set func(entry *types.Entry) error

//Helper database access layer for accessing entries
type Helper struct {
	db *sql.DB
}

//AddDuckEntry add a duck entry to the database
func (h *Helper) AddDuckEntry(entry *types.Entry) error {
	query := `
		INSERT INTO duck_entries (
			id,
			fed_time,
			food,
			kind_of_food,
			amount_of_food,
			location,
			number_of_ducks
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := h.db.Exec(
		query,
		entry.ID,
		entry.TimeFed,
		entry.Food.Name,
		entry.Food.Kind,
		entry.AmountOfFood,
		entry.Location,
		entry.NumberOfDucks,
	)

	return err
}

package data

import (
	"database/sql"

	"github.com/bmanth60/DuckTracker/types"
)

type Set func(entry *types.Entry) error

type Helper struct {
	db *sql.DB
}

func (h *Helper) GetDuckEntries() (*types.Entry, error) {
	query := `
		SELECT
			id,
			fed_time,
			food,
			kind_of_food,
			amount_of_food,
			location,
			number_of_ducks
		FROM duck_entries
	`

	result := new(types.Entry)
	err := h.db.QueryRow(query).Scan(
		&result.ID,
		&result.TimeFed,
		&result.Food.Name,
		&result.Food.Kind,
		&result.AmountOfFood,
		&result.Location,
		&result.NumberOfDucks,
	)

	return result, err
}

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

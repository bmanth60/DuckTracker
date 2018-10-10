package data

import (
	"database/sql"

	"github.com/bmanth60/DuckTracker/types"
)

type Helper struct {
	db *sql.DB
}

func (h *Helper) GetDucks() (*types.Entry, error) {
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

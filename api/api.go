package api

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bmanth60/DuckTracker/data"
	dterr "github.com/bmanth60/DuckTracker/errors"
	"github.com/bmanth60/DuckTracker/types"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type API int

func (a *API) HandleForm(form url.Values, add data.Set) []string {
	var err error
	messages := []string{}
	result := new(types.Entry)

	// Extract values from form
	result.NumberOfDucks, err = a.ValidateNumber(form.Get("num_ducks"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "number of ducks").Error())
	}

	result.TimeFed, err = a.ValidateDate(form.Get("time_fed"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "time fed").Error())
	}

	result.Location, err = a.ValidateText(form.Get("location"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "location").Error())
	}

	result.AmountOfFood, err = a.ValidateNumber(form.Get("food_amount"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "amount of food").Error())
	}

	result.Food.Name, err = a.ValidateText(form.Get("food_name"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "name of food").Error())
	}

	result.Food.Kind, err = a.ValidateText(form.Get("food_kind"))
	if err != nil {
		messages = append(messages, errors.Wrap(err, "kind of food").Error())
	}

	// If values do not pass validation, return error list
	if len(messages) != 0 {
		return messages
	}

	// Insert data into database
	result.ID = xid.New().String()
	err = add(result)
	if err != nil {
		return []string{err.Error()}
	}

	return nil
}

func (a *API) ValidateDate(value string) (time.Time, error) {
	format := "2006-01-02T15:04:05" //ISO8601 format
	return time.Parse(format, value)
}

func (a *API) ValidateNumber(value string) (int, error) {
	if value == "" {
		return -1, errors.Wrap(dterr.ErrInvalidValue, "value cannot be empty")
	}

	return strconv.Atoi(value)
}

func (a *API) ValidateText(value string) (string, error) {
	if value == "" {
		return "", errors.Wrap(dterr.ErrInvalidValue, "value cannot be empty")
	}

	if len(value) > 50 {
		return "", errors.Wrap(dterr.ErrInvalidValue, "value cannot exceed 50 characters")
	}

	return value, nil
}

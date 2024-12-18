package main

import (
	"errors"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerPowerOutagesTime(w http.ResponseWriter, r *http.Request) {
	type response struct {
		PowerOutagesTime string `json:"power_outages_time"`
	}

	weekday := time.Now().Weekday()
	powerOutagesTime, err := schedulePowerOutagesTime(weekday.String())
	if err != nil {
		write500Error(w)
		return
	}

	writeJSONResponse(w, 200, response{
		PowerOutagesTime: powerOutagesTime,
	})
}

func schedulePowerOutagesTime(weekday string) (string, error) {
	switch weekday {
	case "Monday":
		return "From 08:30 AM to 12:00 PM", nil
	case "Tuesday":
		return "From 1:00 PM to 3:00 PM", nil
	case "Wednesday":
		return "From 2:00 PM to 4:00 PM", nil
	case "Thursday":
		return "From 3:00 PM to 5:00 PM", nil
	case "Friday":
		return "From 4:00 PM to 6:00 PM", nil
	case "Saturday":
		return "From 5:00 PM to 7:00 PM", nil
	case "Sunday":
		return "From 6:00 PM to 8:00 PM", nil
	default:
		return "", errors.New("invalid day of week")
	}
}

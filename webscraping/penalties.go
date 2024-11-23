package webscraping

import (
	"encoding/json"
	"math"
	"math/rand"
	"time"
)

type PenaltyType int

const (
	SPEEDING PenaltyType = iota
	WRONGTURN
	PARKING
)

var penalty = map[PenaltyType]string{
	SPEEDING:  "SPEEDING",
	WRONGTURN: "WRONGTURN",
	PARKING:   "PARKING",
}

type Penalty struct {
	IssuedOn string
	Amount   float64
	Reason   string
}

type ResponsePenaltiesObject struct {
	Plate     string
	OwnerId   string
	Penalties []Penalty
}

func FetchPenalties(plate string, ownerId string) string {
	lowerBoundLimit := 1
	upperBoundLimit := 5
	serverProcessingTimeInSec := rand.Intn(upperBoundLimit-lowerBoundLimit) + lowerBoundLimit
	duration := time.Duration(serverProcessingTimeInSec) * time.Second
	ticketsNumber := rand.Intn(10)
	penaltiesList := make([]Penalty, ticketsNumber)
	details := ResponsePenaltiesObject{
		Plate:     plate,
		OwnerId:   ownerId,
		Penalties: penaltiesList,
	}
	for i := 0; i < ticketsNumber; i++ {
		issuanceDays := rand.Intn(365)
		ticketType := rand.Intn(3)
		currentTicket := Penalty{
			IssuedOn: time.Now().AddDate(0, 0, -issuanceDays).Format(time.RFC3339),
			Amount:   math.Round((1 + rand.Float64()) * 100000),
			Reason:   penalty[PenaltyType(ticketType)],
		}
		details.Penalties = append(details.Penalties, currentTicket)
	}
	time.Sleep(duration)
	jsonResponse, jsonErr := json.Marshal(details)
	if jsonErr != nil {
		return ""
	}
	return string(jsonResponse)
}

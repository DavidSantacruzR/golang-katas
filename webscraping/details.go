package webscraping

import (
	"math/rand"
	"time"
)

type InsuranceCompany int

const (
	BOLIVAR InsuranceCompany = iota
	AXA
	HDI
)

var insurer = map[InsuranceCompany]string{
	BOLIVAR: "BOLIVAR",
	AXA:     "AXA",
	HDI:     "HDI",
}

type Insurance struct {
	Issued     string
	Amount     float64
	ValidUntil string
	Insurer    string
}

type ResponseObject struct {
	Plate             string
	OwnerId           string
	Insurance         Insurance
	ValidLicense      bool
	LicenseValidUntil string
	DrivingState      string
}

func FetchVehicleDetails(plate string, ownerId string) ResponseObject {
	lowerBoundLimit := 1
	upperBoundLimit := 5
	serverProcessingTimeInSec := rand.Intn(upperBoundLimit-lowerBoundLimit) + lowerBoundLimit
	duration := time.Duration(serverProcessingTimeInSec) * time.Second
	randomInsuranceValidDate := rand.Intn(100)
	randomInsurer := rand.Intn(3)
	details := ResponseObject{
		Plate:   plate,
		OwnerId: ownerId,
		Insurance: Insurance{
			ValidUntil: time.Now().AddDate(0, 0, -randomInsuranceValidDate).Format(time.RFC3339),
			Issued:     time.Now().AddDate(-1, 0, 0).Format(time.RFC3339),
			Amount:     rand.Float64(),
			Insurer:    insurer[InsuranceCompany(randomInsurer)],
		},
		ValidLicense:      true,
		LicenseValidUntil: time.Now().AddDate(1, 0, 0).Format(time.RFC3339),
		DrivingState:      "CAPITAL_DISTRICT",
	}
	time.Sleep(duration)
	return details
}

package main

import (
	"fmt"
	"golang/webscraping"
)

func main() {
	vehicleDetails := webscraping.FetchVehicleDetails("ABC123", "123456789")
	vehiclePenalties := webscraping.FetchPenalties("ABC123", "123456789")
	fmt.Printf("\n%s\n", vehicleDetails)
	fmt.Printf("\n%s\n", vehiclePenalties)
	webscraping.RunApi()
}

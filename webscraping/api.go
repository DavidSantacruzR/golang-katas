package webscraping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VehicleRequest struct {
	Plate   string
	OwnerId string
}

type VehicleResponse struct {
	Details   any
	Penalties any
}

type Result struct {
	Field string
	Value any
}

func RunApi() {
	router := gin.Default()
	router.GET("/healthCheck", healthCheck)
	router.POST("/vehicleData", fetchVehicleData)
	_ = router.Run()
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}

func fetchVehicleData(c *gin.Context) {
	var body VehicleRequest
	if err := c.BindJSON(&body); err != nil {
		return
	}
	resultChannel := make(chan Result)
	response := VehicleResponse{}
	go func() {
		vehicleDetails := FetchVehicleDetails(body.Plate, body.OwnerId)
		resultChannel <- Result{Field: "Details", Value: vehicleDetails}
	}()
	go func() {
		penaltiesDetails := FetchPenalties(body.Plate, body.OwnerId)
		resultChannel <- Result{Field: "Penalties", Value: penaltiesDetails}
	}()
	for i := 0; i < 2; i++ {
		result := <-resultChannel
		switch result.Field {
		case "Details":
			response.Details = result.Value
		case "Penalties":
			response.Penalties = result.Value
		}
	}
	c.JSON(http.StatusOK, response)
}

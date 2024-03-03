package services

import (
	"github.com/nurmeden/assisted-team/models"
)

func GetDXBtoBKK(airFareSearchResponse models.AirFareSearchResponse) ([]models.PricedItinerary, int) {
	var flights []models.PricedItinerary
	count := 0
	for _, v := range airFareSearchResponse.PricedItineraries {
		destination := ""
		source := v.OnwardPricedItinerary.Flights[0].Source
		for _, flight := range v.OnwardPricedItinerary.Flights {
			destination = flight.Destination
		}
		if source == "DXB" && destination == "BKK" {
			count++
			flights = append(flights, v)
		}
	}

	return flights, count
}

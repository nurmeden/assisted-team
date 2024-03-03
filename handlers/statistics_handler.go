package handlers

import (
	"encoding/json"
	"github.com/nurmeden/assisted-team/models"
	"github.com/nurmeden/assisted-team/services"
	"log"
	"net/http"
)

func GetStatistics(airFareSearchResponse models.AirFareSearchResponse) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		result := services.GetCheapAndExpensiveFlights2(airFareSearchResponse)

		if err := json.NewEncoder(rw).Encode(result); err != nil {
			log.Println("failed to send response: ", err)
			return
		}
	}
}

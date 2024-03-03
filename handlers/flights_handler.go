package handlers

import (
	"encoding/json"
	"github.com/nurmeden/assisted-team/models"
	"github.com/nurmeden/assisted-team/services"
	"log"
	"net/http"
)

const xmlFile1 = "RS_Via-3.xml"

func GetAllFlightsDXBtoBKK(airFareSearchResponse models.AirFareSearchResponse) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		result := make(map[string]any)

		flights, count := services.GetDXBtoBKK(airFareSearchResponse)

		result["flights"] = flights
		result["count"] = count

		if err := json.NewEncoder(rw).Encode(result); err != nil {
			log.Println("failed to send response: ", err)
			return
		}
	}
}

package services

import (
	"fmt"
	"github.com/nurmeden/assisted-team/models"
	"math"
	"time"
)

const (
	cheapFlight     = "cheap_flight"
	expensiveFlight = "expensive_flight"
	typeDate        = "2006-01-02T1504"
	fastFlight      = "fast_flight"
	longFlight      = "long_flight"
)

func GetCheapAndExpensiveFlights2(airFareSearchResponse models.AirFareSearchResponse) map[string]any {
	result := make(map[string]any)
	var maxAmount float64
	minAmount := math.MaxFloat64
	result[cheapFlight] = maxAmount
	result[expensiveFlight] = minAmount
	for _, v := range airFareSearchResponse.PricedItineraries {
		if result[cheapFlight].(float64) < v.Pricing.ServiceCharges[2].Amount {
			result[cheapFlight] = maxAmount
		}
		if result[expensiveFlight].(float64) > v.Pricing.ServiceCharges[2].Amount {
			result[expensiveFlight] = maxAmount
		}

		for _, flight := range v.OnwardPricedItinerary.Flights {
			dateArrival, _ := time.Parse(typeDate, flight.ArrivalTimeStamp)
			dateDeparture, _ := time.Parse(typeDate, flight.DepartureTimeStamp)
			duration := dateArrival.Sub(dateDeparture)

			if _, ok := result[fastFlight]; !ok {
				result[fastFlight] = make(map[string]interface{})
			}

			if val, ok := result[fastFlight].(map[string]any)["duration"]; !ok || val == nil {
				result = setDurationAndFlight(result, duration, v, fastFlight)
			} else if dur, ok := val.(time.Duration); ok && dur > duration {
				result = setDurationAndFlight(result, duration, v, longFlight)
			}

			if _, ok := result[longFlight]; !ok {
				result[longFlight] = make(map[string]interface{})
			}

			if val, ok := result[longFlight].(map[string]any)["duration"]; !ok || val == nil {
				result = setDurationAndFlight(result, duration, v, longFlight)
			} else if dur, ok := val.(time.Duration); ok && dur < duration {
				result = setDurationAndFlight(result, duration, v, fastFlight)
			}
		}
	}
	fmt.Println(result)
	return result
}

func setDurationAndFlight(result map[string]any, duration time.Duration, flight models.PricedItinerary, typeFlight string) map[string]any {
	result[typeFlight].(map[string]interface{})["duration"] = duration
	result[typeFlight].(map[string]interface{})["flight"] = flight.OnwardPricedItinerary
	return result
}

func getAmount(serviceCharge models.ServiceCharge, typeAmount string, prevAmount float64, result *map[string]any, flightIt models.FlightItinerary) (*map[string]any, float64) {
	switch typeAmount {
	case "max":
		if prevAmount < serviceCharge.Amount {
			prevAmount = serviceCharge.Amount
			(*result)[cheapFlight] = flightIt
		}
	case "min":
		if prevAmount > serviceCharge.Amount {
			prevAmount = serviceCharge.Amount
			(*result)[expensiveFlight] = flightIt
		}
	}
	return result, prevAmount
}

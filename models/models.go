package models

import "encoding/xml"

type TypeFlight struct {
	fastFlight      string `default:"fast_flight"`
	longFlight      string `default:"long_flight"`
	cheapFlight     string `default:"cheap_flight"`
	expensiveFlight string `default:"expensive_flight"`
}

type AirFareSearchResponse struct {
	XMLName           xml.Name          `xml:"AirFareSearchResponse"`
	RequestTime       string            `xml:"RequestTime,attr"`
	ResponseTime      string            `xml:"ResponseTime,attr"`
	RequestId         string            `xml:"RequestId"`
	PricedItineraries []PricedItinerary `xml:"PricedItineraries>Flights"`
}

type PricedItinerary struct {
	OnwardPricedItinerary FlightItinerary `xml:"OnwardPricedItinerary>Flights"`
	ReturnPricedItinerary FlightItinerary `xml:"ReturnPricedItinerary>Flights"`
	Pricing               Pricing         `xml:"Pricing"`
}

type FlightItinerary struct {
	Flights []Flight `xml:"Flight"`
}

type Flight struct {
	Carrier            Carrier `xml:"Carrier"`
	FlightNumber       string  `xml:"FlightNumber"`
	Source             string  `xml:"Source"`
	Destination        string  `xml:"Destination"`
	DepartureTimeStamp string  `xml:"DepartureTimeStamp"`
	ArrivalTimeStamp   string  `xml:"ArrivalTimeStamp"`
	Class              string  `xml:"Class"`
	NumberOfStops      int     `xml:"NumberOfStops"`
	FareBasis          string  `xml:"FareBasis"`
	WarningText        string  `xml:"WarningText"`
	TicketType         string  `xml:"TicketType"`
}

type Carrier struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Pricing struct {
	Currency       string          `xml:"currency,attr"`
	ServiceCharges []ServiceCharge `xml:"ServiceCharges"`
}

type ServiceCharge struct {
	Type       string  `xml:"type,attr"`
	ChargeType string  `xml:"ChargeType,attr"`
	Amount     float64 `xml:",chardata"`
}

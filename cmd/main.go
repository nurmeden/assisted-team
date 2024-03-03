package main

import (
	"github.com/nurmeden/assisted-team/handlers"
	"github.com/nurmeden/assisted-team/models"
	"github.com/nurmeden/assisted-team/utils"
	"log"
	"net/http"
)

const (
	xmlFile1 = "xml/RS_Via-3.xml"
	xmlFIle2 = "xml/RS_ViaOW.xml"
)

func main() {
	var airFareSearchResponse models.AirFareSearchResponse
	decodeXML, err := utils.DecodeXML(xmlFile1, airFareSearchResponse)
	if err != nil {
		log.Fatalf("Error decoding XML: %v", err)
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("/getAllFlightsToDXBtoBKK", handlers.GetAllFlightsDXBtoBKK(*decodeXML))
	router.HandleFunc("/statistics", handlers.GetStatistics(*decodeXML))
	log.Fatal(http.ListenAndServe(":8080", router))
}

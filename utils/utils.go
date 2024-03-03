package utils

import (
	"encoding/xml"
	"fmt"
	"github.com/nurmeden/assisted-team/models"
	"os"
)

func DecodeXML(xmlFilePath string, airFareSearchResponse models.AirFareSearchResponse) (*models.AirFareSearchResponse, error) {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening XML file: %v", err)
	}

	defer func(xmlFile *os.File) {
		if closeErr := xmlFile.Close(); closeErr != nil {
			fmt.Printf("Error closing XML file: %v\n", closeErr)
		}
	}(xmlFile)

	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&airFareSearchResponse); err != nil {
		return nil, fmt.Errorf("error decoding XML: %v", err)
	}

	return &airFareSearchResponse, nil
}

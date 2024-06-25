// utils/utils.go

package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"ga/flights" // Substitua "seu-usuario" pelo seu usuário ou nome do módulo
)

// ReadFlights lê os dados de voos de um arquivo CSV
func ReadFlights(fileName string) ([]flights.Flight, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var flightsData []flights.Flight

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Parse each record into Flight struct
		cost, _ := strconv.Atoi(record[4])
		flight := flights.Flight{
			Origin:        record[0],
			Destination:   record[1],
			DepartureTime: record[2],
			ArrivalTime:   record[3],
			Cost:          cost,
		}

		flightsData = append(flightsData, flight)
	}

	return flightsData, nil
}

// OrganizeFlights organiza os voos com base na origem e no destino
func OrganizeFlights(allFlights []flights.Flight) map[string][]flights.Flight {
	flightsOrganized := map[string][]flights.Flight{
		"from_LIS": nil, "to_LIS": nil,
		"from_MAD": nil, "to_MAD": nil,
		"from_CDG": nil, "to_CDG": nil,
		"from_DUB": nil, "to_DUB": nil,
		"from_BRU": nil, "to_BRU": nil,
		"from_LHR": nil, "to_LHR": nil,
	}

	for _, flight := range allFlights {
		switch flight.Origin {
		case "LIS":
			flightsOrganized["from_LIS"] = append(flightsOrganized["from_LIS"], flight)
		case "MAD":
			flightsOrganized["from_MAD"] = append(flightsOrganized["from_MAD"], flight)
		case "CDG":
			flightsOrganized["from_CDG"] = append(flightsOrganized["from_CDG"], flight)
		case "DUB":
			flightsOrganized["from_DUB"] = append(flightsOrganized["from_DUB"], flight)
		case "BRU":
			flightsOrganized["from_BRU"] = append(flightsOrganized["from_BRU"], flight)
		case "LHR":
			flightsOrganized["from_LHR"] = append(flightsOrganized["from_LHR"], flight)
		}

		switch flight.Destination {
		case "LIS":
			flightsOrganized["to_LIS"] = append(flightsOrganized["to_LIS"], flight)
		case "MAD":
			flightsOrganized["to_MAD"] = append(flightsOrganized["to_MAD"], flight)
		case "CDG":
			flightsOrganized["to_CDG"] = append(flightsOrganized["to_CDG"], flight)
		case "DUB":
			flightsOrganized["to_DUB"] = append(flightsOrganized["to_DUB"], flight)
		case "BRU":
			flightsOrganized["to_BRU"] = append(flightsOrganized["to_BRU"], flight)
		case "LHR":
			flightsOrganized["to_LHR"] = append(flightsOrganized["to_LHR"], flight)
		}
	}

	return flightsOrganized
}

// GetLatestArrivalFlight retorna o voo com o horário de chegada mais tarde
func GetLatestArrivalFlight(fs []flights.Flight) flights.Flight {
	if len(fs) == 0 {
		return flights.Flight{}
	}

	latestFlight := fs[0]
	layout := "15:04"

	for _, flight := range fs[1:] {
		latestTime, _ := time.Parse(layout, latestFlight.ArrivalTime)
		currentTime, _ := time.Parse(layout, flight.ArrivalTime)

		if currentTime.After(latestTime) {
			latestFlight = flight
		}
	}

	return latestFlight
}

// GetFirstDepartureFlight retorna o voo com o horário de partida mais cedo
func GetFirstDepartureFlight(fs []flights.Flight) flights.Flight {
	if len(fs) == 0 {
		return flights.Flight{}
	}

	firstFlight := fs[0]
	layout := "15:04"

	for _, flight := range fs[1:] {
		firstTime, _ := time.Parse(layout, firstFlight.DepartureTime)
		currentTime, _ := time.Parse(layout, flight.DepartureTime)

		if currentTime.Before(firstTime) {
			firstFlight = flight
		}
	}

	return firstFlight
}

// IndexesToOrganizedFlightList retorna a chave para acessar a lista de voos organizada com base nos índices
func IndexesToOrganizedFlightList(toOrFromIndex, researcherIndex int) string {
	researchers := []string{"to_LIS", "to_MAD", "to_CDG", "to_DUB", "to_BRU", "to_LHR"}
	researchersFrom := []string{"from_LIS", "from_MAD", "from_CDG", "from_DUB", "from_BRU", "from_LHR"}

	if toOrFromIndex == 0 {
		return researchers[researcherIndex]
	} else if toOrFromIndex == 1 {
		return researchersFrom[researcherIndex]
	}

	return ""
}

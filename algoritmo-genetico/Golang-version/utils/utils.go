// utils/utils.go

package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"ga/flights" // Substitua "seu-usuario" pelo seu usuário ou nome do modulo
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

// OrganizeFlights organiza os voos com base na origem e no destino MELHOR
func OrganizeFlights(allFlights []flights.Flight) map[string]map[string][]flights.Flight {
	flightsOrganized := map[string]map[string][]flights.Flight{
		"from": {
			"LIS": nil,
			"MAD": nil,
			"CDG": nil,
			"DUB": nil,
			"BRU": nil,
			"LHR": nil,
		},
		"to": {
			"LIS": nil,
			"MAD": nil,
			"CDG": nil,
			"DUB": nil,
			"BRU": nil,
			"LHR": nil,
		},
	}

	for _, flight := range allFlights {
		// Going (to the events) flights
		switch flight.Origin {
		case "LIS":
			flightsOrganized["from"]["LIS"] = append(flightsOrganized["from"]["LIS"], flight)
		case "MAD":
			flightsOrganized["from"]["MAD"] = append(flightsOrganized["from"]["MAD"], flight)
		case "CDG":
			flightsOrganized["from"]["CDG"] = append(flightsOrganized["from"]["CDG"], flight)
		case "DUB":
			flightsOrganized["from"]["DUB"] = append(flightsOrganized["from"]["DUB"], flight)
		case "BRU":
			flightsOrganized["from"]["BRU"] = append(flightsOrganized["from"]["BRU"], flight)
		case "LHR":
			flightsOrganized["from"]["LHR"] = append(flightsOrganized["from"]["LHR"], flight)
		}

		// Returning (from the events) flights
		switch flight.Destination {
		case "LIS":
			flightsOrganized["to"]["LIS"] = append(flightsOrganized["to"]["LIS"], flight)
		case "MAD":
			flightsOrganized["to"]["MAD"] = append(flightsOrganized["to"]["MAD"], flight)
		case "CDG":
			flightsOrganized["to"]["CDG"] = append(flightsOrganized["to"]["CDG"], flight)
		case "DUB":
			flightsOrganized["to"]["DUB"] = append(flightsOrganized["to"]["DUB"], flight)
		case "BRU":
			flightsOrganized["to"]["BRU"] = append(flightsOrganized["to"]["BRU"], flight)
		case "LHR":
			flightsOrganized["to"]["LHR"] = append(flightsOrganized["to"]["LHR"], flight)
		}
	}

	return flightsOrganized
}

func IndexToAirportCode(idx int) string {
	airportCodeMap := map[int]string{
		0: "LIS",
		1: "MAD",
		2: "CDG",
		3: "DUB",
		4: "BRU",
		5: "LHR",
	}

	airportCode, ok := airportCodeMap[idx]
	if !ok {
		panic("IndexToAirportCode: index out of map bounds!")
	}

	return airportCode
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

func TotalIndividualFlightCosts(individual []flights.Flight) int {
	if len(individual) != 12 {
		panic("TotalFlightCosts: more flights provided than needed for a singular individual")
	}

	total := 0
	for _, flight := range individual {
		total += flight.Cost
	}

	return total
}

// TotalAirportWaitingMinutes retorna o tempo total de espera nos aeroportos
func TotalAirportWaitingMinutes(individual []flights.Flight) int {
	if len(individual) != 12 {
		panic("TotalAirportWaitingMinutes: more flights provided than needed for a singular individual")
	}

	goingFlights := []flights.Flight{}
	returningFlights := []flights.Flight{}

	// Voos de ida (from)
	for i := 0; i < 12; i += 2 {
		goingFlights = append(goingFlights, individual[i])
	}
	latestArrivalFlight := GetLatestArrivalFlight(goingFlights)

	// Voos de volta (to)
	for v := 1; v < 12; v += 2 {
		returningFlights = append(returningFlights, individual[v])
	}
	firstDepartureFlight := GetFirstDepartureFlight(returningFlights)

	totalWaiting := 0
	latestArrivalTime, _ := time.Parse("15:04", latestArrivalFlight.ArrivalTime)
	firstDepartureTime, _ := time.Parse("15:04", firstDepartureFlight.DepartureTime)

	// Voos de ida (from)
	for i := 0; i < 12; i += 2 {
		arrivalTime, _ := time.Parse("15:04", individual[i].ArrivalTime)
		totalWaiting += int(latestArrivalTime.Sub(arrivalTime).Minutes())
		// fmt.Println("(ida)", arrivalTime.Format("15:04"), "->", latestArrivalTime.Format("15:04"), "=", int(latestArrivalTime.Sub(arrivalTime).Minutes()))
	}

	// Voos de volta (to)
	for v := 1; v < 12; v += 2 {
		departureTime, _ := time.Parse("15:04", individual[v].DepartureTime)
		totalWaiting += int(departureTime.Sub(firstDepartureTime).Minutes())
		// fmt.Println("(vta)", firstDepartureTime.Format("15:04"), "->", departureTime.Format("15:04"), "=", int(departureTime.Sub(firstDepartureTime).Minutes()))
	}

	return totalWaiting
}

// OpenCSV abre e lida com algumas questões do arquivo CSV criado.
func OpenCSV(filename string) (*os.File, *csv.Writer) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo CSV: %v", err)
	}

	writer := csv.NewWriter(file)

	header := []string{"Best", "Average", "Worst"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalf("Erro ao escrever o cabecalho no arquivo CSV: %v", err)
	}

	return file, writer
}

func FinalResult(flights []flights.Flight) {
	fmt.Println("Tempo medio de espera por pesquisador:", TotalAirportWaitingMinutes(flights)/6)
	fmt.Println("Tempo de espera total:", TotalAirportWaitingMinutes(flights))
	fmt.Println("Custo medio de passagem por pesquisador:", TotalIndividualFlightCosts(flights)/6)
	fmt.Println("Custo total:", TotalIndividualFlightCosts(flights))
}

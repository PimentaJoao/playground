package individuals

import (
	"fmt"
	"math/rand"
	"time"

	"ga/flights"
	"ga/researchers"
	"ga/utils"
)

// Individual representa um indivíduo com voos para cada pesquisador
type Individual struct {
	Researchers []researchers.Researcher
}

// NewIndividual cria um novo indivíduo com voos aleatórios para cada pesquisador
func NewIndividual(flightsOrganized map[string][]flights.Flight) Individual {
	return Individual{
		Researchers: []researchers.Researcher{
			{
				GoingFlight:     flightsOrganized["from_LIS"][rand.Intn(len(flightsOrganized["from_LIS"]))],
				ReturningFlight: flightsOrganized["to_LIS"][rand.Intn(len(flightsOrganized["to_LIS"]))],
			},
			{
				GoingFlight:     flightsOrganized["from_MAD"][rand.Intn(len(flightsOrganized["from_MAD"]))],
				ReturningFlight: flightsOrganized["to_MAD"][rand.Intn(len(flightsOrganized["to_MAD"]))],
			},
			{
				GoingFlight:     flightsOrganized["from_CDG"][rand.Intn(len(flightsOrganized["from_CDG"]))],
				ReturningFlight: flightsOrganized["to_CDG"][rand.Intn(len(flightsOrganized["to_CDG"]))],
			},
			{
				GoingFlight:     flightsOrganized["from_DUB"][rand.Intn(len(flightsOrganized["from_DUB"]))],
				ReturningFlight: flightsOrganized["to_DUB"][rand.Intn(len(flightsOrganized["to_DUB"]))],
			},
			{
				GoingFlight:     flightsOrganized["from_BRU"][rand.Intn(len(flightsOrganized["from_BRU"]))],
				ReturningFlight: flightsOrganized["to_BRU"][rand.Intn(len(flightsOrganized["to_BRU"]))],
			},
			{
				GoingFlight:     flightsOrganized["from_LHR"][rand.Intn(len(flightsOrganized["from_LHR"]))],
				ReturningFlight: flightsOrganized["to_LHR"][rand.Intn(len(flightsOrganized["to_LHR"]))],
			},
		},
	}
}

// TotalFlightCosts retorna o custo total dos voos de todos os pesquisadores
func (ind *Individual) TotalFlightCosts() int {
	totalCost := 0
	for _, researcher := range ind.Researchers {
		totalCost += researcher.GoingFlight.Cost + researcher.ReturningFlight.Cost
	}
	return totalCost
}

// TotalAirportWaitingMinutes retorna o tempo total de espera nos aeroportos
func (ind *Individual) TotalAirportWaitingMinutes() int {
	goingFlights := []flights.Flight{
		ind.Researchers[0].GoingFlight,
		ind.Researchers[1].GoingFlight,
		ind.Researchers[2].GoingFlight,
		ind.Researchers[3].GoingFlight,
		ind.Researchers[4].GoingFlight,
		ind.Researchers[5].GoingFlight,
	}

	returningFlights := []flights.Flight{
		ind.Researchers[0].ReturningFlight,
		ind.Researchers[1].ReturningFlight,
		ind.Researchers[2].ReturningFlight,
		ind.Researchers[3].ReturningFlight,
		ind.Researchers[4].ReturningFlight,
		ind.Researchers[5].ReturningFlight,
	}

	latestArrivalFlight := utils.GetLatestArrivalFlight(goingFlights)
	firstDepartureFlight := utils.GetFirstDepartureFlight(returningFlights)

	totalWaiting := 0
	latestTime, _ := time.Parse("15:04", latestArrivalFlight.ArrivalTime)
	firstTime, _ := time.Parse("15:04", firstDepartureFlight.DepartureTime)

	for _, researcher := range ind.Researchers {
		arrivalTime, _ := time.Parse("15:04", researcher.GoingFlight.ArrivalTime)
		totalWaiting += int(latestTime.Sub(arrivalTime).Minutes())

		departureTime, _ := time.Parse("15:04", researcher.ReturningFlight.DepartureTime)
		totalWaiting += int(departureTime.Sub(firstTime).Minutes())
	}

	return totalWaiting
}

// IndexToFlight retorna o nome correto da chave para acessar a lista de voos de interesse
func (ind *Individual) IndexToFlight(toOrFromIndex, researcherIndex int) string {
	if toOrFromIndex == 0 {
		switch researcherIndex {
		case 0:
			return "to_LIS"
		case 1:
			return "to_MAD"
		case 2:
			return "to_CDG"
		case 3:
			return "to_DUB"
		case 4:
			return "to_BRU"
		case 5:
			return "to_LHR"
		}
	} else if toOrFromIndex == 1 {
		switch researcherIndex {
		case 0:
			return "from_LIS"
		case 1:
			return "from_MAD"
		case 2:
			return "from_CDG"
		case 3:
			return "from_DUB"
		case 4:
			return "from_BRU"
		case 5:
			return "from_LHR"
		}
	}

	return ""
}

// String retorna uma representação em string do indivíduo
func (ind *Individual) String() string {
	LISCost := ind.Researchers[0].GoingFlight.Cost + ind.Researchers[0].ReturningFlight.Cost
	MADCost := ind.Researchers[1].GoingFlight.Cost + ind.Researchers[1].ReturningFlight.Cost
	CDGCost := ind.Researchers[2].GoingFlight.Cost + ind.Researchers[2].ReturningFlight.Cost
	DUBCost := ind.Researchers[3].GoingFlight.Cost + ind.Researchers[3].ReturningFlight.Cost
	BRUCost := ind.Researchers[4].GoingFlight.Cost + ind.Researchers[4].ReturningFlight.Cost
	LHRCost := ind.Researchers[5].GoingFlight.Cost + ind.Researchers[5].ReturningFlight.Cost

	return fmt.Sprintf("LIS researcher flights:\n  Cost: %d\nMAD researcher flights:\n  Cost: %d\nCDG researcher flights:\n  Cost: %d\nDUB researcher flights:\n  Cost: %d\nBRU researcher flights:\n  Cost: %d\nLHR researcher flights:\n  Cost: %d\n",
		LISCost, MADCost, CDGCost, DUBCost, BRUCost, LHRCost)
}

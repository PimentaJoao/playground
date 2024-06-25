package flights

import "fmt"

// Flight representa um voo
type Flight struct {
	Origin        string
	Destination   string
	DepartureTime string
	ArrivalTime   string
	Cost          int
}

// NewFlight cria uma nova instância de Flight
func NewFlight(origin, destination, departureTime, arrivalTime string, cost int) Flight {
	return Flight{
		Origin:        origin,
		Destination:   destination,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		Cost:          cost,
	}
}

// String retorna uma representação em string do voo
func (f Flight) String() string {
	return fmt.Sprintf("From: %s, To: %s, Departure: %s, Arrival: %s, Cost: %d",
		f.Origin, f.Destination, f.DepartureTime, f.ArrivalTime, f.Cost)
}

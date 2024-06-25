package researchers

import (
	"fmt"
	"ga/flights" // Substitua "seu-usuario" pelo seu usuário ou nome do módulo
)

// Researcher representa um pesquisador
type Researcher struct {
	GoingFlight     flights.Flight
	ReturningFlight flights.Flight
}

// NewResearcher cria uma nova instância de Researcher
func NewResearcher(goingFlight, returningFlight flights.Flight) *Researcher {
	return &Researcher{
		GoingFlight:     goingFlight,
		ReturningFlight: returningFlight,
	}
}

// String retorna uma representação em string do pesquisador
func (r *Researcher) String() string {
	return fmt.Sprintf("Researcher:\n  Going: %s\n  Returning: %s",
		r.GoingFlight.String(), r.ReturningFlight.String())
}

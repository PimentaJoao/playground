package algoritmo_genetico

import (
	"fmt"
	"math/rand"
	"sort"

	"ga/flights"
	"ga/utils"
)

// GA struct representa o algoritmo genetico
type GA struct {
	populationSize int
	generations    int
	tournamentN    float64
	tournamentK    float64
	pCrossover     float64
	pMutation      float64
	elitism        bool

	population []flights.Flight
}

// NewGA cria uma nova instancia do algoritmo genetico
func NewGA(populationSize, generations int, tournamentN, tournamentK, pCrossover, pMutation float64, elitism bool) GA {
	return GA{
		populationSize: populationSize,
		generations:    generations,
		tournamentN:    tournamentN,
		tournamentK:    tournamentK,
		pCrossover:     pCrossover,
		pMutation:      pMutation,
		elitism:        elitism,
		population:     make([]flights.Flight, 0),
	}
}

// InitPopulation inicializa a populacao inicial de individuos
func (ga *GA) InitPopulation(flightsOrganized map[string]map[string][]flights.Flight) {

	// Itera a nivel "individuo"
	for i := 0; i < ga.populationSize; i++ {

		// Itera a nivel "pesquisador", adicionando seus voos de ida (from) e de volta (to)
		for j := 0; j < 6; j++ {
			airport := utils.IndexToAirportCode(j)

			numberOfGoingFlights := len(flightsOrganized["from"][airport])
			randomGoingFlightIndex := rand.Intn(numberOfGoingFlights)
			goingFlight := flightsOrganized["from"][airport][randomGoingFlightIndex]

			numberOfReturningFlights := len(flightsOrganized["to"][airport])
			randomReturningFlightIndex := rand.Intn(numberOfReturningFlights)
			returningFlight := flightsOrganized["to"][airport][randomReturningFlightIndex]

			ga.population = append(ga.population, goingFlight, returningFlight)
		}
	}
}

// ShowPopulation exibe todos os individuos na populacao
func (ga *GA) ShowPopulation() {
	for i, flight := range ga.population {
		fmt.Println(flight.String())

		if i%12 == 0 && i != 0 {
			fmt.Println("")
			fmt.Println("")
		}
	}
}

// GetRankedPopulation devolve todos os individuos da populacao ordenados em ordem de fitness crescente
func (ga *GA) GetRankedPopulation() []flights.Flight {
	return ga.rankPopulation(ga.population)
}

// ShowPopulationWithFitness exibe todos os individuos na populacao e seu fitness individual
func (ga *GA) ShowPopulationWithFitness() {
	for i, flight := range ga.population {
		if i%12 == 0 && i != 0 {
			fmt.Println("")
			fmt.Println("")
		}

		if i%12 == 0 {
			fmt.Println(ga.Fitness(ga.population[i : i+12]))
		}

		fmt.Println(flight.String())

	}
}

// ShowPopulationWithFitness exibe todos os individuos na populacao e seu fitness individual
func (ga *GA) ShowPopulationFitness() {
	for i := 0; i < len(ga.population); i++ {
		if i%12 == 0 {
			fmt.Println("individuo", i/12, ":", ga.Fitness(ga.population[i:i+12]))
		}
	}
}

// ShowRankedPopulation exibe todos os individuos da populacao ordenados por ordem crescente de fitness
func (ga *GA) ShowRankedPopulation() {
	rankedPopulation := ga.rankPopulation(ga.population)

	for i := 0; i < len(rankedPopulation); i++ {
		if i%12 == 0 {
			fmt.Println("individuo", i/12, ":", ga.Fitness(rankedPopulation[i:i+12]))
		}
	}
}

// GetBestAverageWorstFitness retorna a melhor, media e pior aptidao da populacao
func (ga *GA) GetBestAverageWorstFitness() (best, avg, worst int) {
	rankedPopulation := ga.rankPopulation(ga.population)

	best = ga.Fitness(rankedPopulation[:12])

	total := 0
	for i := 0; i < len(rankedPopulation); i++ {
		if i%12 == 0 {
			total += ga.Fitness(rankedPopulation[i : i+12])
		}
	}
	avg = total / ga.populationSize

	worst = ga.Fitness(rankedPopulation[len(rankedPopulation)-12:])

	return best, avg, worst
}

// Selection realiza a selecao dos individuos para a proxima geracao
func (ga *GA) Selection(flightsOrganized map[string]map[string][]flights.Flight) {
	newPopulation := make([]flights.Flight, 0)

	// Elitismo
	if ga.elitism {
		newPopulation = append(newPopulation, ga.rankPopulation(ga.population)[:12]...)
	}

	// Torneio
	for len(newPopulation) < ga.populationSize*12 {
		newPopulation = append(newPopulation, ga.tournament()...)
	}

	// Mutacao
	mutatedPopulation := ga.applyMutation(flightsOrganized, newPopulation)

	// Crossover
	crossedOverPopulation := ga.applyCrossover(mutatedPopulation)

	ga.population = crossedOverPopulation
}

// tournament realiza um torneio para selecionar um individuo
func (ga *GA) tournament() []flights.Flight {
	tournamentSize := int(float64(ga.populationSize) * ga.tournamentN)
	if tournamentSize < 2 {
		tournamentSize = 2
	}

	tournamentIndexes := make(map[int]bool)
	for len(tournamentIndexes) < tournamentSize {
		index := rand.Intn(ga.populationSize)
		tournamentIndexes[index] = true
	}

	var tournamentPopulation []flights.Flight
	for index := range tournamentIndexes {
		tournamentPopulation = append(tournamentPopulation, ga.population[index*12:index*12+12]...)
	}

	rankedTournamentPopulation := ga.rankPopulation(tournamentPopulation)

	tournamentR := rand.Float64()
	if tournamentR < ga.tournamentK {
		return rankedTournamentPopulation[:12]
	}
	return rankedTournamentPopulation[len(rankedTournamentPopulation)-12:]
}

// applyMutation aplica mutacao na populacao
func (ga *GA) applyMutation(flightsOrganized map[string]map[string][]flights.Flight, population []flights.Flight) []flights.Flight {

	if len(population)%12 != 0 {
		panic("applyMutation: population length not in correct size for full individuals")
	}

	// Trata do elitismo
	start := 0
	if ga.elitism {
		start = 12
	}

	for i := start; i < len(population); i += 12 {
		// So realiza mutacao no individuo se mutationChance for menor que o parametro
		mutationChance := rand.Intn(100)
		if mutationChance > int(ga.pMutation*100.0) {
			continue
		}

		randomFlightIdx := rand.Intn(12)

		direction := ""
		airport := ""

		if randomFlightIdx%2 == 0 {
			direction = "from"
			airport = utils.IndexToAirportCode(randomFlightIdx / 2)
		} else {
			direction = "to"
			airport = utils.IndexToAirportCode((randomFlightIdx - 1) / 2)
		}

		newFlightIdx := rand.Intn(len(flightsOrganized[direction][airport]))

		population[i+randomFlightIdx] = flightsOrganized[direction][airport][newFlightIdx]
	}

	return population
}

func (ga *GA) applyCrossover(population []flights.Flight) []flights.Flight {
	if len(population)%12 != 0 {
		panic("applyCrossover: population length not in correct size for full individuals")
	}

	type individualAndCrossoverStatus struct {
		flights     []flights.Flight
		crossedOver bool
	}
	iacs := []individualAndCrossoverStatus{}

	// Populando lista auxiliar para lidar com crossover
	for i := 0; i < len(population); i++ {
		if i%12 == 0 {
			iacs = append(iacs, individualAndCrossoverStatus{
				flights:     population[i : i+12],
				crossedOver: false,
			})
		}
	}

	// Trata do elitismo
	start := 0
	if ga.elitism {
		start = 1
	}
	iacs[0].crossedOver = true

	for i, iac := range iacs[start:] {
		// So realiza crossover no individuo se crossoverChance for menor que o parametro
		crossoverChance := rand.Intn(100)
		if crossoverChance > int(ga.pCrossover*100.0) {
			continue
		}

		// Seleciona com qual individuo o crossover vai acontecer
		// Roda ate encontrar um elemento que ainda nao sofreu crossover
		otherIndividualIdx := rand.Intn(len(iacs[start:]))
		hasCheckedEveryOne := 0
		for iacs[start:][otherIndividualIdx].crossedOver || otherIndividualIdx == i {
			otherIndividualIdx = rand.Intn(len(iacs[start:]))
			hasCheckedEveryOne++

			if hasCheckedEveryOne == len(iacs[start:]) {
				break
			}
		}

		thisIndividual := iac
		otherIndividual := iacs[start:][otherIndividualIdx]

		// Ponto de crossover [1..5]
		crossoverPoint := (rand.Intn(6-1) + 1) * 2

		// Coleta cabeca e calda dos genes deste individuo
		thisIndividualHead := thisIndividual.flights[:crossoverPoint]
		thisIndividualTail := append([]flights.Flight{}, thisIndividual.flights[crossoverPoint:]...)

		// Coleta cabeca e calda dos genes do segundo individuo
		otherIndividualHead := otherIndividual.flights[:crossoverPoint]
		otherIndividualTail := append([]flights.Flight{}, otherIndividual.flights[crossoverPoint:]...)

		// Cria novos individuos filhos misturados (crossover)
		newIndividual1 := []flights.Flight{}
		newIndividual1 = append(newIndividual1, thisIndividualHead...)
		newIndividual1 = append(newIndividual1, otherIndividualTail...)

		newIndividual2 := []flights.Flight{}
		newIndividual2 = append(newIndividual2, otherIndividualHead...)
		newIndividual2 = append(newIndividual2, thisIndividualTail...)

		// DEBUG [applyCrossover]: descomentar caso seja necessária uma análise mais profunda

		// fmt.Println("CROSSOVER POINT:", crossoverPoint)
		// fmt.Println("i:", i)
		// fmt.Println("otherIndividualIdx:", otherIndividualIdx)
		// fmt.Println("Antes:")
		// fmt.Println("")
		// for j := 0; j < 12; j++ {
		// 	fmt.Println(iacs[i+start].flights[j].String())
		// }
		// fmt.Println("")
		// for j := 0; j < 12; j++ {
		// 	fmt.Println(iacs[otherIndividualIdx+start].flights[j].String())
		// }
		// fmt.Println("")
		// fmt.Println("Depois:")
		// fmt.Println("")
		// for j := 0; j < 12; j++ {
		// 	fmt.Println(newIndividual1[j].String())
		// }
		// fmt.Println("")
		// for j := 0; j < 12; j++ {
		// 	fmt.Println(newIndividual2[j].String())
		// }
		// fmt.Println("")
		// fmt.Println("Cabecas:")
		// fmt.Println("")
		// for j := 0; j < len(thisIndividualHead); j++ {
		// 	fmt.Println(thisIndividualHead[j].String())
		// }
		// fmt.Println("")
		// for j := 0; j < len(otherIndividualHead); j++ {
		// 	fmt.Println(otherIndividualHead[j].String())
		// }
		// fmt.Println("")
		// fmt.Println("Caudas:")
		// fmt.Println("")
		// for j := 0; j < len(thisIndividualTail); j++ {
		// 	fmt.Println(thisIndividualTail[j].String())
		// }
		// fmt.Println("")
		// for j := 0; j < len(otherIndividualTail); j++ {
		// 	fmt.Println(otherIndividualTail[j].String())
		// }
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("")

		// Substitui pais pelos seus filhos misturados
		iacs[i+start].flights = newIndividual1
		iacs[i+start].crossedOver = true
		iacs[otherIndividualIdx+start].flights = newIndividual2
		iacs[otherIndividualIdx+start].crossedOver = true
	}

	crossedOverPopulation := []flights.Flight{}
	for _, iac := range iacs {
		crossedOverPopulation = append(crossedOverPopulation, iac.flights...)
	}

	return crossedOverPopulation
}

// Fitness calcula a aptidao de um individuo (12 voos)
func (ga *GA) Fitness(individual []flights.Flight) int {
	if len(individual) != 12 {
		panic("Fitness: more flights provided than needed for a singular individual")
	}

	// return utils.TotalIndividualFlightCosts(individual) + utils.TotalAirportWaitingMinutes(individual)
	return utils.TotalIndividualFlightCosts(individual)
}

// rankPopulation classifica a populacao com base no fitness score
func (ga *GA) rankPopulation(population []flights.Flight) []flights.Flight {
	if len(population)%12 != 0 {
		panic("rankPopulation: population length not in correct size for full individuals")
	}

	type individualAndFitness struct {
		flights []flights.Flight
		fitness int
	}

	iafs := []individualAndFitness{}

	for i := 0; i < len(population); i++ {
		if i%12 == 0 {
			iafs = append(iafs, individualAndFitness{
				flights: population[i : i+12],
				fitness: ga.Fitness(population[i : i+12]),
			})
		}
	}

	// Ordenar os chunks por fitness score (em ordem crescente)
	sort.Slice(iafs, func(i, j int) bool {
		return iafs[i].fitness < iafs[j].fitness
	})

	rankedPopulation := []flights.Flight{}
	for _, iaf := range iafs {
		rankedPopulation = append(rankedPopulation, iaf.flights...)
	}

	return rankedPopulation
}

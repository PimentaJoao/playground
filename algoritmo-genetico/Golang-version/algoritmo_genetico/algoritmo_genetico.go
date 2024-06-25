package algoritmo_genetico

import (
	"fmt"
	"math/rand"
	"sort"

	"ga/flights"
	"ga/individuals"
)

// GA struct representa o algoritmo genético
type GA struct {
	populationSize int
	generations    int
	tournamentN    float64
	tournamentK    float64
	pCrossover     float64
	pMutation      float64
	elitism        bool

	population []individuals.Individual
}

// NewGA cria uma nova instância do algoritmo genético
func NewGA(populationSize, generations int, tournamentN, tournamentK, pCrossover, pMutation float64, elitism bool) GA {
	return GA{
		populationSize: populationSize,
		generations:    generations,
		tournamentN:    tournamentN,
		tournamentK:    tournamentK,
		pCrossover:     pCrossover,
		pMutation:      pMutation,
		elitism:        elitism,
		population:     make([]individuals.Individual, 0),
	}
}

// InitPopulation inicializa a população inicial de indivíduos
func (ga *GA) InitPopulation(flightsOrganized map[string][]flights.Flight) {
	for i := 0; i < ga.populationSize; i++ {
		ga.population = append(ga.population, individuals.NewIndividual(flightsOrganized))
	}
}

// ShowPopulation exibe todos os indivíduos na população
func (ga *GA) ShowPopulation() {
	for _, individual := range ga.population {
		fmt.Println(individual)
	}
}

// GetPopulation retorna a população atual
func (ga *GA) GetPopulation() []individuals.Individual {
	return ga.population
}

// ShowPopulationFitness exibe a aptidão de todos os indivíduos na população
func (ga *GA) ShowPopulationFitness() {
	for _, individual := range ga.population {
		fmt.Println(ga.Fitness(individual))
	}
}

// GetBestAverageWorstFitness retorna a melhor, média e pior aptidão da população
func (ga *GA) GetBestAverageWorstFitness() (best, avg, worst int) {
	rankedPopulation := ga.rankPopulation(ga.population)

	best = ga.Fitness(rankedPopulation[0])

	total := 0
	for _, individual := range rankedPopulation {
		total += ga.Fitness(individual)
	}
	avg = total / len(rankedPopulation)

	worst = ga.Fitness(rankedPopulation[len(rankedPopulation)-1])

	return best, avg, worst
}

// Selection realiza a seleção dos indivíduos para a próxima geração
func (ga *GA) Selection(flightsOrganized map[string][]flights.Flight) {
	newPopulation := make([]individuals.Individual, 0)

	// Elitismo
	if ga.elitism {
		newPopulation = append(newPopulation, ga.rankPopulation(ga.population)[0])
	}

	// Torneio
	for len(newPopulation) < ga.populationSize {
		newPopulation = append(newPopulation, ga.tournament())
	}

	// Mutação
	mutatedPopulation := ga.applyMutation(flightsOrganized, newPopulation)

	// Crossover
	// crossedOverPopulation := ga.applyCrossover(mutatedPopulation)

	ga.population = mutatedPopulation
}

// tournament realiza um torneio para selecionar um indivíduo
func (ga *GA) tournament() individuals.Individual {
	tournamentSize := int(float64(ga.populationSize) * ga.tournamentN)
	if tournamentSize < 2 {
		tournamentSize = 2
	}

	tournamentIndexes := make(map[int]bool)
	for len(tournamentIndexes) < tournamentSize {
		index := rand.Intn(ga.populationSize)
		tournamentIndexes[index] = true
	}

	var tournamentPopulation []individuals.Individual
	for index := range tournamentIndexes {
		tournamentPopulation = append(tournamentPopulation, ga.population[index])
	}

	rankedTournamentPopulation := ga.rankPopulation(tournamentPopulation)

	tournamentR := rand.Float64()
	if tournamentR < ga.tournamentK {
		return rankedTournamentPopulation[0]
	}
	return rankedTournamentPopulation[len(rankedTournamentPopulation)-1]
}

// applyMutation aplica mutação na população
func (ga *GA) applyMutation(flightsOrganized map[string][]flights.Flight, population []individuals.Individual) []individuals.Individual {

	// Trata do elitismo
	start := 0
	if ga.elitism {
		start = 1
	}

	for i := start; i < len(population); i++ {
		// Só realiza mutação no indivíduo se mutationChance for menor que o parâmetro
		mutationChance := rand.Intn(100)
		if mutationChance > int(ga.pMutation*100.0) {
			continue
		}

		population[i] = individuals.NewIndividual(flightsOrganized)
	}

	return population
}

func (ga *GA) applyCrossover(population []individuals.Individual) []individuals.Individual {

	// Trata do elitismo
	start := 0
	if ga.elitism {
		start = 1
	}

	for i := start; i < len(population)-1; i++ {
		// Só realiza crossover no indivíduo se crossoverChance for menor que o parâmetro
		crossoverChance := rand.Intn(100)
		if crossoverChance > int(ga.pCrossover*100.0) {
			continue
		}

		researcherIdx := rand.Intn(6)

		population[i].Researchers[researcherIdx], population[i+1].Researchers[researcherIdx] =
			population[i+1].Researchers[researcherIdx], population[i].Researchers[researcherIdx]
	}

	return population
}

// fitness calcula a aptidão de um indivíduo
func (ga *GA) Fitness(individual individuals.Individual) int {
	return individual.TotalFlightCosts() + individual.TotalAirportWaitingMinutes()
}

// rankPopulation classifica a população com base na aptidão
func (ga *GA) rankPopulation(population []individuals.Individual) []individuals.Individual {
	// Ordena os indivíduos com base na pontuação de aptidão.
	// Supõe que uma pontuação de aptidão menor é melhor e deve estar no início da lista.
	return SortByFitness(population, ga.Fitness)
}

// SortByFitness ordena um slice de Individual pelo menor score de fitness
func SortByFitness(population []individuals.Individual, fitnessFunc func(individuals.Individual) int) []individuals.Individual {
	sorter := &individualSorter{
		population:  population,
		fitnessFunc: fitnessFunc,
	}

	sort.Stable(sorter)

	return sorter.population
}

// individualSorter implementa a interface sort.Interface para ordenar Individuals por fitness
type individualSorter struct {
	population  []individuals.Individual
	fitnessFunc func(individuals.Individual) int
}

func (s individualSorter) Len() int {
	return len(s.population)
}

func (s individualSorter) Less(i, j int) bool {
	return s.fitnessFunc(s.population[i]) < s.fitnessFunc(s.population[j])
}

func (s individualSorter) Swap(i, j int) {
	s.population[i], s.population[j] = s.population[j], s.population[i]
}

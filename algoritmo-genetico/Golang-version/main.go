package main

import (
	"fmt"
	"ga/algoritmo_genetico"
	"ga/utils"
	"strconv"
)

func main() {
	flights, err := utils.ReadFlights("flights.txt")
	if err != nil {
		fmt.Println("Error reading flights:", err)
		return
	}

	organizedFlights := utils.OrganizeFlights(flights)

	/*
		Como usar organizedFlights:

			organizedFlights["<direcao>"]["<codigo do aeroporto>"]

		exemplo:

			allFlightsGoingToLisbon = organizedFlights["to"]["LIS"]
	*/

	populationSize := 200
	generations := 1000
	tournamentN := 0.20
	tournamentK := 0.75
	pCrossover := 0.8
	pMutation := 0.3
	elitism := true

	ga := algoritmo_genetico.NewGA(populationSize, generations, tournamentN, tournamentK, pCrossover, pMutation, elitism)

	ga.InitPopulation(organizedFlights)

	/*
		Alguns métodos foram criados para monitorar o progresso do GA, por exemplo:

		ga.ShowPopulation()
		exibe os voos dos indivíduos de toda a população

		ga.ShowPopulationWithFitness()
		exibe os voos dos indivíduos de toda a população + seu fitness

		ga.ShowPopulationFitness()
		exibe o fitness de cada indivíduo da população completa

		ga.ShowRankedPopulation()
		exibe um ranking de todos os indivíduos dado seu fitness
	*/

	file, writer := utils.OpenCSV("dados.csv")
	defer file.Close()
	defer writer.Flush()

	for i := 0; i < generations; i++ {
		ga.Selection(organizedFlights)
		best, avg, worst := ga.GetBestAverageWorstFitness()
		if i%100 == 0 {
			fmt.Println("best", "avg", "worst")
			fmt.Println(best, avg, worst)
		}
		data := []string{strconv.Itoa(best), strconv.Itoa(avg), strconv.Itoa(worst)}
		writer.Write(data)
	}

	bestFlights := ga.GetRankedPopulation()[:12]
	for _, flight := range bestFlights {
		fmt.Println(flight.String())
	}

	fmt.Println()

	utils.FinalResult(bestFlights)
}

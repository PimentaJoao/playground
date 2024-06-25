package main

import (
	"encoding/csv"
	"fmt"
	"ga/algoritmo_genetico"
	"ga/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	flights, err := utils.ReadFlights("flights.txt")
	if err != nil {
		fmt.Println("Error reading flights:", err)
		return
	}

	organizedFlights := utils.OrganizeFlights(flights)

	populationSize := 200
	generations := 1000
	tournamentN := 0.20
	tournamentK := 0.75
	pCrossover := 0.8
	pMutation := 0.3
	elitism := true

	ga := algoritmo_genetico.NewGA(populationSize, generations, tournamentN, tournamentK, pCrossover, pMutation, elitism)

	ga.InitPopulation(organizedFlights)

	best, avg, worst := ga.GetBestAverageWorstFitness()
	fmt.Println("best, avg, worst:")
	fmt.Println(best, avg, worst)

	filename := "dados.csv"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Best", "Average", "Worst"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalf("Erro ao escrever o cabeçalho no arquivo CSV: %v", err)
	}

	for i := 0; i < generations; i++ {
		ga.Selection(organizedFlights)
		best, avg, worst = ga.GetBestAverageWorstFitness()
		data := []string{strconv.Itoa(best), strconv.Itoa(avg), strconv.Itoa(worst)}
		err = writer.Write(data)
		if err != nil {
			log.Fatalf("Erro ao escrever os dados no arquivo CSV: %v", err)
		}
	}

	best, avg, worst = ga.GetBestAverageWorstFitness()
	fmt.Println("best, avg, worst:")
	fmt.Println(best, avg, worst)

	if best < 3500 {
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[0].Researchers[0].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[0].Researchers[0].ReturningFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[1].Researchers[1].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[1].Researchers[1].ReturningFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[2].Researchers[2].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[2].Researchers[2].ReturningFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[3].Researchers[3].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[3].Researchers[3].ReturningFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[4].Researchers[4].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[4].Researchers[4].ReturningFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[5].Researchers[5].GoingFlight)
		fmt.Println(algoritmo_genetico.SortByFitness(ga.GetPopulation(), ga.Fitness)[5].Researchers[5].ReturningFlight)
	}
}

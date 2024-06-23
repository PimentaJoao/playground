import utils
from GA import GA
from datetime import datetime

def main():
    file_name = "flights.txt"

    all_flights = utils.read_flights(file_name)

    flights_organized = utils.organize_flights(all_flights)

    # Exemplo, pegar todos os voos de Roma até Lisboa:
    #    for flight in flights_organized["to_LIS"]:
    #        print(flight)

    population_size = 3
    generations = 10
    tournament_N = 0.2 # % of population in tournament
    tournament_K = 0.75 # odds of better individual to succeed
    p_crossover = 0.8
    p_mutation = 0.05
    elitism = True

    ga = GA(population_size, generations, tournament_N, tournament_K, p_crossover, p_mutation, elitism)

    ga.initPopulation(flights_organized)

    # DEBUG
    individual = ga.getPopulation()
    idx = 1
    for researcher in individual[0].researchers:
        print(f"Pesquisador {idx}:")
        print(f"  {researcher.going_flight}")
        print(f"  {researcher.returning_flight}")
        idx = idx + 1

    ga.showPopulationFitness()

    # print("")
    # print("Aleatorio")
    # print("best, avg, worst:")
    # print(ga.getBestAverageWorstFitness())
    # print("")
    # ga.selection(flights_organized)

    # print("1a selecao")
    # print("best, avg, worst:")
    # print(ga.getBestAverageWorstFitness())
    # print("")

    # ga.selection(flights_organized)
    # print("2a selecao")
    # print("best, avg, worst:")
    # print(ga.getBestAverageWorstFitness())
    # print("")

    # ga.selection(flights_organized)
    # print("3a selecao")
    # print("best, avg, worst:")
    # print(ga.getBestAverageWorstFitness())
    # print("")

    # for _ in range(27):
    #     ga.selection(flights_organized)
    # print("30a selecao")
    # print("best, avg, worst:")
    # print(ga.getBestAverageWorstFitness())
    # print("")


if __name__ == "__main__":
    main()

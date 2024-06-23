from individuals import Individual
from random import uniform, randrange, choice
import utils

class GA:
    def __init__(self, population_size, generations, tourn_N, tourn_K, p_crossover, p_mutation, elitism):
        self.population_size = population_size
        self.generations = generations
        self.tournament_N = tourn_N
        self.tournament_K = tourn_K
        self.p_crossover = p_crossover
        self.p_mutation = p_mutation
        self.elitism = elitism

        self.population = []
        self.current_generation = 1

    def initPopulation(self, flights_organized):
        for _ in range(self.population_size):
            self.population.append(Individual(flights_organized))
   
    def showPopulation(self):
        for individual in self.population:
            print(individual)

    def getPopulation(self):
        return self.population

    def showPopulationFitness(self):
        for individual in self.population:
            print(self._fitness(individual))

    def getBestAverageWorstFitness(self):
        ranked_population = self._rankPopulation(self.population)

        best = self._fitness(ranked_population[0])

        total = 0
        for individual in ranked_population:
            total = total + self._fitness(individual)
        avg = total / len(ranked_population)

        worst = self._fitness(ranked_population[len(ranked_population)-1])

        return best, avg, worst

    def selection(self, flights_organized):
        new_population = []

        # elitism
        if self.elitism == True:
            new_population.append(self._rankPopulation(self.population)[0])

        # tournament
        while len(new_population) < self.population_size:
            new_population.append(self._tournament())
        self.population = new_population

        # mutation
        # self._applyMutation(flights_organized)



    def _tournament(self) -> Individual:
        tournament_size = (int) (self.population_size * self.tournament_N)

        # Creates a list of unique indexes, mapping which individuals are going
        # to compose the tournament
        tournament_population_indexes = set()
        while len(tournament_population_indexes) < tournament_size:
            tournament_population_indexes.add(randrange(self.population_size))

        # Grabs the selected individuals, based on their indexes
        tournament_population = []
        for index in tournament_population_indexes:
            tournament_population.append(self.population[index])
        
        ranked_tournament_population = self._rankPopulation(tournament_population)

        tournament_R = uniform(0, 1)

        if tournament_R < self.tournament_K:
            # best individual
            return ranked_tournament_population[0]
        else:
            # worst individual
            return ranked_tournament_population[len(ranked_tournament_population)-1]

    def _applyCrossover(self):
        # TODO: Implement.
        return
    
    def _applyMutation(self, flights_organized):
        
        for individual_index in range(len(self.population)):
            mutation_chance = uniform(0, 1)

            if mutation_chance > self.p_mutation:
                # won't mutate; continue to next individual
                continue

            # 0 = "to"   (returning to country of origin flight)
            # 1 = "from" (going to country of event)
            to_or_from = randrange(1)

            # which "bit" (researcher) to mutate [0..5]
            researcher_index = randrange(6) 

            if to_or_from == 0: # mutate returning flight
                flight_list_name = utils.indexes_to_organized_flight_list(0, researcher_index)

                new_returning_flight = choice(flights_organized[flight_list_name])
                self.population[individual_index].researchers[researcher_index].returning_flight = new_returning_flight

            if to_or_from == 1: # mutate going flight
                flight_list_name = utils.indexes_to_organized_flight_list(1, researcher_index)

                new_going_flight = choice(flights_organized[flight_list_name])
                self.population[individual_index].researchers[researcher_index].going_flight = new_going_flight
                
        return

    def _fitness(self, individual: Individual):
        return individual.totalFlightCosts() + individual.totalAirportWaitingMinutes()
    
    def _rankPopulation(self, population:list):
        # Orders individuals based on fitness score.
        # Assumes smaller fitness score = better = first on the list.
        return sorted(population, key=lambda x: self._fitness(x))

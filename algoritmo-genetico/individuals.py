from researchers import Researcher
from random import choice
import utils
from datetime import datetime

class Individual:

    # Initializes individual with random departure and returning flights for each researcher
    def __init__(self, flights_organized):
        self.researchers = [
            Researcher(choice(flights_organized["from_LIS"]), choice(flights_organized["to_LIS"])),
            Researcher(choice(flights_organized["from_MAD"]), choice(flights_organized["to_MAD"])),
            Researcher(choice(flights_organized["from_CDG"]), choice(flights_organized["to_CDG"])),
            Researcher(choice(flights_organized["from_DUB"]), choice(flights_organized["to_DUB"])),
            Researcher(choice(flights_organized["from_BRU"]), choice(flights_organized["to_BRU"])),
            Researcher(choice(flights_organized["from_LHR"]), choice(flights_organized["to_LHR"]))
        ]

    def totalFlightCosts(self):
        LIS_cost = self.researchers[0].going_flight.cost + self.researchers[0].returning_flight.cost
        MAD_cost = self.researchers[1].going_flight.cost + self.researchers[1].returning_flight.cost
        CDG_cost = self.researchers[2].going_flight.cost + self.researchers[2].returning_flight.cost
        DUB_cost = self.researchers[3].going_flight.cost + self.researchers[3].returning_flight.cost
        BRU_cost = self.researchers[4].going_flight.cost + self.researchers[4].returning_flight.cost
        LHR_cost = self.researchers[5].going_flight.cost + self.researchers[5].returning_flight.cost
        return LIS_cost + MAD_cost + CDG_cost + DUB_cost + BRU_cost + LHR_cost
    
    def totalAirportWaitingMinutes(self) -> int:
        latest_arrival_flight = utils.getLatestArrivalFlight([
            self.researchers[0].going_flight,
            self.researchers[1].going_flight,
            self.researchers[2].going_flight,
            self.researchers[3].going_flight,
            self.researchers[4].going_flight,
            self.researchers[5].going_flight
        ])

        first_departure_flight = utils.getFirstDepartureFlight([
            self.researchers[0].returning_flight,
            self.researchers[1].returning_flight,
            self.researchers[2].returning_flight,
            self.researchers[3].returning_flight,
            self.researchers[4].returning_flight,
            self.researchers[5].returning_flight
        ])

        total_waiting = 0 # in minutes

        for researcher in self.researchers:
            # waiting on arrival
            latest_time = datetime.strptime(latest_arrival_flight.arrival_time, "%H:%M")
            arrival_time = datetime.strptime(researcher.going_flight.arrival_time, "%H:%M")
            total_waiting = total_waiting + int((latest_time - arrival_time).total_seconds() / 60)

            # waiting for departure
            first_time = datetime.strptime(first_departure_flight.departure_time, "%H:%M")
            departure_time = datetime.strptime(researcher.going_flight.departure_time, "%H:%M")
            total_waiting = total_waiting + int((departure_time - first_time).total_seconds() / 60)

        return total_waiting

    # Returns the correct string for accessing the flight list of interest, based on which researcher and on whether 
    # the departure flight or the returning flight is required ("toOrFromIndex")
    def indexToFlight(self, toOrFromIndex, researcherIndex):

        # 0 = "to" origin country flight (returning flight)
        if toOrFromIndex == 0:
            if researcherIndex == 0:
                return "to_LIS"
            if researcherIndex == 1:
                return "to_MAD"
            if researcherIndex == 2:
                return "to_CDG"
            if researcherIndex == 3:
                return "to_DUB"
            if researcherIndex == 4:
                return "to_BRU"
            if researcherIndex == 5:
                return "to_LHR"
            
        # 1 = "from" origin country flight (departure flight)
        if toOrFromIndex == 1:
            if researcherIndex == 0:
                return "to_LIS"
            if researcherIndex == 1:
                return "to_MAD"
            if researcherIndex == 2:
                return "to_CDG"
            if researcherIndex == 3:
                return "to_DUB"
            if researcherIndex == 4:
                return "to_BRU"
            if researcherIndex == 5:
                return "to_LHR"

    def __str__(self):
        LIS_cost = self.researchers[0].going_flight.cost + self.researchers[0].returning_flight.cost
        MAD_cost = self.researchers[1].going_flight.cost + self.researchers[1].returning_flight.cost
        CDG_cost = self.researchers[2].going_flight.cost + self.researchers[2].returning_flight.cost
        DUB_cost = self.researchers[3].going_flight.cost + self.researchers[3].returning_flight.cost
        BRU_cost = self.researchers[4].going_flight.cost + self.researchers[4].returning_flight.cost
        LHR_cost = self.researchers[5].going_flight.cost + self.researchers[5].returning_flight.cost
        
        return f"LIS researcher flights:\n  Cost: {LIS_cost}\nMAD researcher flights:\n  Cost: {MAD_cost}\nCDG researcher flights:\n  Cost: {CDG_cost}\nDUB researcher flights:\n  Cost: {DUB_cost}\nBRU researcher flights:\n  Cost: {BRU_cost}\nLHR researcher flights:\n  Cost: {LHR_cost}\n"
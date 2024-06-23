from flight import Flight
from typing import List
from datetime import datetime

# reads flights txt file
def read_flights(file_name):
    flights = []

    with open(file_name, 'r') as file:
        for line in file:
            data = line.strip().split(',')

            # Airports
            origin = data[0]
            destination = data[1]

            # Arrival and departure times
            departure_time = data[2]
            arrival_time = data[3]

            # Monetary cost
            cost = int(data[4])

            # Adds flight entry
            flights.append(Flight(origin, destination, departure_time, arrival_time, cost))

    return flights

# organizes flights based on if they are leaving the researchers country of origin (ex: from_LIS) or
# if they are returning to the researchers country of origin (ex: to_LIS)
def organize_flights(all_flights):
    flights_organized = {
        "from_LIS": [],
        "to_LIS": [],

        "from_MAD": [],
        "to_MAD": [],

        "from_CDG": [],
        "to_CDG": [],

        "from_DUB": [],
        "to_DUB": [],

        "from_BRU": [],
        "to_BRU": [],

        "from_LHR": [],
        "to_LHR": []
    }

    for flight in all_flights:
        if flight.origin == "LIS":
            flights_organized["from_LIS"].append(flight)
        if flight.destination == "LIS":
            flights_organized["to_LIS"].append(flight)

        if flight.origin == "MAD":
            flights_organized["from_MAD"].append(flight)
        if flight.destination == "MAD":
            flights_organized["to_MAD"].append(flight)

        if flight.origin == "CDG":
            flights_organized["from_CDG"].append(flight)
        if flight.destination == "CDG":
            flights_organized["to_CDG"].append(flight)

        if flight.origin == "DUB":
            flights_organized["from_DUB"].append(flight)
        if flight.destination == "DUB":
            flights_organized["to_DUB"].append(flight)

        if flight.origin == "BRU":
            flights_organized["from_BRU"].append(flight)
        if flight.destination == "BRU":
            flights_organized["to_BRU"].append(flight)

        if flight.origin == "LHR":
            flights_organized["from_LHR"].append(flight)
        if flight.destination == "LHR":
            flights_organized["to_LHR"].append(flight)

    return flights_organized

# Compares list of flight items and reveals the latest one that arrived based on it's arrival time.
def getLatestArrivalFlight(flights: List[Flight]) -> Flight:

    # initialize with first element for further comparisons
    latest_flight = flights[0]

    for flight in flights[1:]:
        latest_flight_arrival_time = datetime.strptime(latest_flight.arrival_time, "%H:%M")
        arrival_time = datetime.strptime(flight.arrival_time, "%H:%M")

        if latest_flight_arrival_time < arrival_time:
            latest_flight = flight


    return latest_flight

# Compares list of flight items and reveals the first one to depart based on it's departure time.
def getFirstDepartureFlight(flights: List[Flight]) -> Flight:

    # initialize with first element for further comparisons
    first_flight = flights[0]

    for flight in flights[1:]:
        first_flight_departure_time = datetime.strptime(first_flight.departure_time, "%H:%M")
        departure_time = datetime.strptime(flight.departure_time, "%H:%M")

        if first_flight_departure_time > departure_time:
            first_flight = flight


    return first_flight

# Returns the correct string for accessing the flight list of interest, based on which researcher [0..5] and on whether 
# the departure flight or the returning flight [0 or 1] is required
def indexes_to_organized_flight_list(to_or_from_index, researcher_index) -> str:

    # 0 = "to" (returning flight)
    if to_or_from_index == 0:
        if researcher_index == 0:
            return "to_LIS"
        if researcher_index == 1:
            return "to_MAD"
        if researcher_index == 2:
            return "to_CDG"
        if researcher_index == 3:
            return "to_DUB"
        if researcher_index == 4:
            return "to_BRU"
        if researcher_index == 5:
            return "to_LHR"
        
    # 1 = "from" (departure flight)
    if to_or_from_index == 1:
        if researcher_index == 0:
            return "from_LIS"
        if researcher_index == 1:
            return "from_MAD"
        if researcher_index == 2:
            return "from_CDG"
        if researcher_index == 3:
            return "from_DUB"
        if researcher_index == 4:
            return "from_BRU"
        if researcher_index == 5:
            return "from_LHR"

            
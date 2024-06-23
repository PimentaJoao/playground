class Flight:
    def __init__(self, origin, destination, departure_time, arrival_time, cost):
        self.origin = origin
        self.destination = destination
        self.departure_time = departure_time
        self.arrival_time = arrival_time
        self.cost = cost

    def __str__(self):
        return f"From: {self.origin}, To: {self.destination}, Departure: {self.departure_time}, Arrival: {self.arrival_time}, Cost: {self.cost}"

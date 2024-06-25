from flight import Flight

class Researcher:
    def __init__(self, going_flight: Flight, returning_flight: Flight):
        # From the perspective of the researchers' origin country
        self.going_flight = going_flight         # "going" to the event
        self.returning_flight = returning_flight # "returning" from the event

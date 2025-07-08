## Flight Booking System

### Requirements
1. User should be able to login into system
2. User should be able to search flights based on arrival + destination + date of travel
3. Select flight based on time and preference
4. Book Flight by entering personal details, make payment
5. System should send confirmation notification.

**Actors** : Cutomers, Airline Authorities, Admin

### Core Entities
1. Flight
- Attributes => FlightNo, AirlineCompany, seatCapacity, Seats: List{Seat}, FlightSchedule: List{Schedule}

2. Airline - Name, Flights: List{Flight}

3. Seat - SeatNo, SeatClass

4. FlightSeat extends Seat - Price, Status: open/booked

5. Airport - Name, Location, Flights: List{Flight}

6. Schedule
- Attributes => Flight, StartAirport, DestinationAirport, Date, StartTime, ArrivalTime, Fare, Status: Ontime/ Delayed/ Cancelled

7. User - Name, email, phone, userId

8. FlightBookingSystem
- Attributes => Users: List{User}, Flights: List{Flight}
- Methods =>
    - GetFlightDetails(startLocation, destinationLocation, date)
    - BookFlight(Flight, User)
    - ConfirmBooking(BookingDetails)

9. BookingDetails
- Attributes => FlightSchedule, User, PNR, StartAirport, DestinationAirport, date

10. Payment Module
11. Notification Server



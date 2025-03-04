# Designing a Movie Ticket Booking System like BookMyShow

### Requirements:
1. The system should display all active theatre or active movies wrt to selected city or location as choosen by user.
2. theatre can have multiple theatre halls, each hall have a seating arrangement, and each hall can run one movie at a time.
3. Within a day, one movie can have multiple shows.
4. The system should allow theatre admins to create, update and remove theatre halls, movies, movie shows and seating arrangements of theatre halls.
5. The User should be able to search movie by language, title, release date and city name.
6. On selecting a movie, all active theatres running that movie present in city should be displayed by system.
7. User should be able to select the theatre and book the ticket.
8. System should redirect to seating arrangement where user should be able to select the seat.
9. User should be able to do payment and confirm the booking.
10. The system should handle concurrent bookings and ensure seat availability is updated (vacant and occupied) in real time.
11. The seat should show different types of seats ( normal, premium, etc) with pricing.
12. System should be scalable to handle a large no.of concureent users and bookings.
13. User should be able to cancel the bookings.

### Core Components: 
1. User : Represents people using the system
2. Theatre
3. Movie
4. Show
4. Booking
5. Payment
6. Notification

### Class Definations:
1. User:
- userId
- name
- email
- phone 

2. Theatre: 
- theatreId
- Address
- theatre Name
- total halls

3. Theatre Hall
- hall Id
- hall Name
- total seats

4. Seat
- Row No
- Column No
- SeatType {Regular, Premium}

5. Movie
- Id 
- Name
- description
- duration
- Language
- Genre
- release date
- List<Actors>
- List<Reviews>

6. Show
- ShowId
- MovieId
- theatreId
- hall Id
- startTime
- endTime
- date
- list<show seat>

7. Show Seat
- showId
- seatId
- status { booked, vacant, blocked}

8. Booking
- TicketId
- List<seats>
- Booking Status
- userId
- showId

9. Payment
- TransactionId
- Payment Status
- Amount
- Payment Mode
- payment Date

10. Notification
- message
- status
- Notification Type

![Class Diagram](./ClassDiagram.png)

### Entity Relationship:
1. theatre to theatre Hall => 1:n 
2. theatre Hall to seat => 1:n
3. theatre Hall to show => 1:n
4. show to show seat => 1:n
5. user to booking => 1:n
6. Booking to payment => 1:1
7. movie to show => 1:n

![Usecase Diagram](./usecaseDiagram.png)

### Design Patterns
1. 




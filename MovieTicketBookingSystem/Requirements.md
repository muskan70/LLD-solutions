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

### Actors : 
1. User -> Just searching/ Checking movies
2. Customer -> booked a movie (need authentication)
3. Frontend Officer
4. theatre Admin
5. System 

### Microservices :
1. User 
2. Movie Catalog
3. Booking System
4. Payment System
5. Notification System

![Class Diagram](./ClassDiagram.png)

### Entities
1. User 
2. Movie
3. theatre
4. theatre Hall
5. Seat
6. Show
7. Show Seat -> handle concurrency here
8. Address
9. Booking
10. Payment
11. Notification

### Entity Relationship:
1. theatre to theatre Hall => 1:n 
2. theatre Hall to seat => 1:n
3. theatre Hall to show => 1:n
4. show to show seat => 1:n
5. user to booking => 1:n
6. Booking to payment => 1:1

![Usecase Diagram](./usecaseDiagram.png)




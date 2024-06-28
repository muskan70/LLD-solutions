# Designing a Movie Ticket Booking System like BookMyShow
1. The system should display all active cinemas or active movies wrt to selected city or location as choosen by user.
2. Cinema can have multiple cinema halls, each hall have a seating arrangement, and each hall can run one movie at a time.
3. Within a day, one movie can have multiple shows.
4. The system should allow cinema admins to create, update and remove cinema halls, movies, movie shows and seating arrangements of cinema halls.
5. The User should be able to search movie by language, title, release date and city name.
6. On selecting a movie, all active cinemas running that movie present in city should be displayed by system.
7. User should be able to select the cinema and book the ticket.
8. System should redirect to seating arrangement where user should be able to select the seat.
9. User should be able to do payment and confirm the booking.
10. The system should handle concurrent bookings and ensure seat availability is updated (vacant and occupied) in real time.
11. The seat should show different types of seats ( normal, premium, etc) with pricing.
12. System should be scalable to handle a large no.of concureent users and bookings.
13. User should be able to cancel the bookings.

# Actors : 
1. User -> Just searching/ Checking movies
2. Customer -> booked a movie (need authentication)
3. Frontend Officer
4. Cinema Admin
5. System 

# Microservices :
1. User 
2. Cinema Catalog
3. Booking System
4. Payment System
5. Notification System




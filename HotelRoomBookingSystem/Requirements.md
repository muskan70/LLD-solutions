## Hotel Room Booking System

##### You're building the backend for a hotel room booking system. Each room can be in one of three states:
- AVAILABLE: The room is free to book
- HELD: Temporarily blocked by a customer (for up to 10 minutes)
- BOOKED: Confirmed reservation

##### The system should support the following operations via APIs:
- View available rooms
- Place a hold on a room for a specific user (expires automatically after 10 minutes)
- Confirm the booking of a held room
- Cancel a hold or a booking

##### Key Requirements:
- A room can only be held by one user at a time
- If a user tries to hold or book a room that's already held or booked, it should fail
- Holds should expire automatically after 10 minutes if not confirmed
- Confirmed bookings are final and cannot be overwritten
- The same user can only hold one room at a time

##### API Documentation
1. POST 'http://localhost:8080/api/rooms/create' 
- data '{"roomType":1}'

2. GET 'http://localhost:8080/api/rooms/status?status=1' 

3. POST 'http://localhost:8080/api/booking/create' \
- data '{ "roomId":1, "userId":4 }'

4. POST 'http://localhost:8080/api/booking/confirm' \
- data '{ "bookingId" : 4 }'

5. POST 'http://localhost:8080/api/booking/cancel' \
- data '{ "bookingId" : 4 }'
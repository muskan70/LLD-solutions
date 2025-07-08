## Car Rental System like zoomcar

### Functional Requirements
1. User should be able to login/signup into system.
2. Enter user location from where he want to book the car
3. List down all cars available at location
4. User can book the car, first by adding booking details likepickup and drop Location.
5. Payment Gateway and Notification system

**Actors** : Operators, Customers

### Core Entities
1. Car
- Attributes => Id, Name, Brand, BookingPrice, kmDriven, manufactureDate, VehicleNo, Status: [Running, Booked, Vacant, NotOperational], VehicleType: [Sedan, HatchBack, SUV], Location
- Methods => ReserveVehicle(), UpdateCarDetails()

2. Reservation
- Attributes => user, car, pickupLocation, dropLocation, pickUpDate, dropDate, rentalPrice, bookStatus, PaymentStatus

3. CarReservationSystem
- Attibutes => Cars : List{Car}, Users: List{User}, AvailableLocations: List{Location}, CarLocationMap : Hashmap{Location, Cars}
- Methos =>
    - GetCarByLocation(location)
    - getCarByType(type)
    - BookCar(reservation)
    - ConfirmBooking()

4. User - name, email, paNo, location, licenseNo
5. PaymentGateway
6. NotificationSystem
    

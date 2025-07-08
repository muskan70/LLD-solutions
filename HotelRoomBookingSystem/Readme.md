## Class Definations for Hotel Booking System like MakeMyTrip/Booking.com

### Functional Requirements
1. User should be able to sign into application.
2. Search on basic of location and date.
3. Select type of room based on preference by user
4. Payment support
5. Booking Confirmation and Notification

**Actors** : Operators (Hotel Owners), Customers

### Core Entities
1. Hotel
- Attributes => Name, uniqueHotelId, Address, Rooms: List{Room}, Rating, Reviews: List{Comment}

2. Address
- Attributes => latitude, logitude, zipcode, state, country, address

3. Room
- Attributes => roomId, hotelId, roomType: [Premium, Delux, Single, Double, ...], price, roomStatus: Occupied/ Booked/ Vacant/ UnderMaintenance/ Available

4. RoomBookingDetails
- Attributes => BookingId, user, hotelId, roomId, dateTo, dateFrom, PaymentStatus, BookingStatus

5. HotelManagementSystem
- Attributes => Hotels: Hashmap{Location, Hotels}, Users: List{User}
- Methods =>
    - signIn(user)
    - searchHotel(Location, date)
    - bookRoom(user, roomBookingDetails)
    - notifyUser(user, roomBookingDetails)

6. Comment - Username, email, content

7. User -  name, email, phone, encrypted password



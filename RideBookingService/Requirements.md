### Problem Statement
Create a low-level design for the backend of a ride-booking service like Uber or Lyft. Consider components such as user registration, location tracking, ride matching, fare calculation, and driver allocation.

### Core Components:
1. User
2. Driver
3. Ride
4. Payment

**User**
- Id Auto increment number
- Name string
- Email string
- Phone int(10)
- Location coordinates 

**Driver** 
- Id Auto increment number
- Name string
- Email string
- Phone int(10)
- Location coordinates -> Index
- RatingStar float64
- VehicleType/DriverType
- VehicleNo 

**Ride**
- Id AutoIncrement number
- DriverId int
- UserId int
- StartingLocation coordinates
- DestinationLocation coordinates
- Distance float64
- TimeTaken/DurationInMinutes int
- StartTimestamp time
- EndTimestamp time
- Status (Requested/Started/Ongoing/Completed/Cancelled)
- Fare float64
- Vehicle Details

**Payment**
- RideId int
- TransactionId string
- ModeOfPayment
- Status (Intiated/Completed/Cancelled)
- Timestamp

### Microservices
1. User Registration and Driver Registration
- CreateNewUser/ SignUP
- Login
- LocationUpdate

2. Location Tracking 
- web socket connection to send pings after every 10s

3. RideMatching
- POST /api/v1/ride/create
> RequestParams: { UserId:xx, StartLocation:xx, Destination:xx }<br>
> Response: { RideId: int, Fare: map{ VehicleType:fareAmount } }<br>
> Function: CreateNewRide(userDetails){<br>
>     Step1: Get fastest Route from third party API<br>
>     Step2: Get time and distance of ride ><br>
>     Step3: fare calculation -> strategy pattern for different situations as well as vehicles<br>
> }

- POST /api/v1/ride/searchDrivers
> RequestParams: { RideId:xx, DriverType:xx }<br>
> Response: waiting for a driver to accept the ride / no driver is available<br>
> Function: SearchNearbyDrivers(ride *Ride, driverType){<br>
>     Step1: get all nearby drivers using starting location<br>
>     Step2: Observer pattern  -> Notify drivers <br>
> }

- POST /api/v1/ride/accept
> RequestParams { DriverId:xx, RideId:xx }<br>
> Response : user has been informed<br>
> Function: AcceptRide(driverId, rideId){<br>
>     Step1: update ride table<br>
>     Step2: user will be notify <br>
> }

- GET /api/v1/ride/details?rideId=xx
> Response: ride details
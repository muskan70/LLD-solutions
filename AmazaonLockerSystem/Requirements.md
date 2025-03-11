# Amazon Locker Facility System

### Requirements
- User should be able to get nearest Locker location.
- User can choose locker facility for package delivery or package return.
- Locker should be optimally assigned based on size of order package i.e SMALL, MEDIUM, LARGE, EXTRA_LARGE
- One time code should be given for opening the locker to both user and delivery Executive.

### Additional Requirements
- Frozen Lockers for frozen food
- How to detect unclaimed packages/items

### Core Components
1. User
2. Order
3. Locker
4. Locker System
5. DeliveryExecutive

### Class Definations
1. User 
- Attributes => UserId, name, email, phone, password, location
- Methods    => OptLockerForDelivery(orderId), OptLockerForReturn(orderId), GetOrderFromLocker(lockerId, code), ReturnOrderToLocker(lockerId, orderId, code)

2. DeliveryExecutive
- Attributes => ExecutiveID, name, email, phone, password
- Methods    => ExecuteDelivery(lockerId, code)

3. Order
- Attributes => OrderId, UserId, DeliveryExecutiveId, PackageSize

4. Locker
- Attributes => LockerId, Location, Size, Code, Status, OrderId, Code, timestamp
- Methods    => GenerateCode(), UpdateStatus()

5. LockerSystem
- Attributes => Location, AvailableLockerMap{Size,Lockers}, ReservedLockerMap{Size,Lockers}
- FindNearestLockerFacility(location), AssignLockerForOrder(orderId), AddOrder(lockerId, orderId, code), RemoveOrder(lockerId, code)

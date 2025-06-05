## LLD-solutions
+ [Amazon Locker System](./AmazaonLockerSystem/Requirements.md)
+ [Appointment Assigner](./AppointmentAssigner/)
+ [Banking System](./BankingSystem/Requirements.md)
+ [Book Manager](./BookManager)
+ [Build in Cache Library](./CacheSystem/InMemoryCacheWithStrategyPattern) : Strategy Design Pattern
+ [Chat Server App](./ChatServerApp/Requirements.md) 
+ [Chess Game](./ChessGame/Requirements.md)
+ [Course Curriculum manager](./CourseCurriculumManager)
+ [Cricket Scoreboard](./CricketScoreboard/Readme.md)
+ [Cron Parser](./CronParser/README.md)
+ [Gaming Score Board System](./GamingScoreBoardSystem)
+ [Gym Centre Session Booking System](./GymSlotBookingManager/Requirements.md)
+ [Hotel Room Booking System](./HotelRoomBookingSystem/Requirements.md)
+ [JobScheduler](./JobScheduler/Requirements.md)
+ [Logging System](./LoggingSystem/README.md)
+ [Meeting Scheduler](./MeetingScheduler/Requirements.md)
+ [Movie Ticket Booking System](./MovieTicketBookingSystem/Requirements.md)
+ [MultiAttribute Key-Value Cache System](./MultiAttributeKeyValueCacheSystem)
+ [N-elevator System](./NElevatorSystem)
+ [NotificationSystem](./NotificationSystem/Requirements.md)
+ [Parking Lot](./ParkingLot)
+ [Publisher Subscriber System](./PublisherSubcriberSystem/Readme.md) : Pull Based Architecture, Observer Design Pattern
+ [Rate Limiter](./RateLimiter/Requirements.md)
+ [Ride Booking Service](./RideBookingService/Requirements.md)
+ [Ride Sharing Service](./RideSharingSystem/Requirements.md)
+ [Snakes and Ladders Game](./SnakesNLaddersGame/Requirements.md) : Factory Design Pattern
+ [Splitwise](./SplitwiseSystem)
+ [Stack Overflow](./StackOverflowSystem)
+ [Tic Tac Toe Game](./TicTacToeGame)
+ [Unix file search API](./UnixFileSearchAPI/README.md) : Composite Design Pattern, Specification Design Pattern 
+ [User Login with JWT Authentication](./UserLogin/README.md)
+ [Vending Machine](./VendingMachineSystem) : State Design Pattern, Singleton Pattern 
+ [Web Scraper with Concurrent Go Routines](./WebScrapingWithConcurrentGoRoutines/Readme.md)

### Steps to run above LLD solutions:
> 1. **create go.mod file** : go mod init {project-name}
> 2. **build code**: go build
> 3. **run code**: ./{project-name}

### LLD Interview Cheat Sheet
1. **Understand the Problem**
- Clarify Requirements: Ask detailed questions to understand the game’s rules, objectives, and constraints.
- Scope Definition: Determine the scope of the design (e.g., core mechanics vs. complete UI/UX design).

2. **Break Down the Problem**
- Identify Core Components: Break down the game into core components (e.g., Player, GameBoard, Dice, etc.).
- Define Interactions: Determine how these components interact with each other.

3. **Class Diagrams**
- Define Classes: Identify the primary classes and their responsibilities.
- Attributes and Methods: Specify key attributes and methods for each class.
- Relationships: Define relationships (associations, aggregations, compositions) between classes.

4. **Design Patterns**
- Identify Relevant Patterns: Utilize design patterns where applicable (e.g., Singleton for game instance, Strategy for different game strategies, Observer for event handling).
- Pattern Application: Clearly articulate why and how each pattern is applied.

5. **Code Structure**
- Modular Design: Ensure that your design promotes modularity and reusability.
- SOLID Principles: Adhere to SOLID principles (Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion).

6. **Implementation Plan**
- Pseudocode : Write pseudocode for critical parts of the system to demonstrate logic.
- Time Management: Allocate specific time slots for different sections (e.g., 15 mins for class design, 10 mins for patterns, 10 mins for pseudocode).

7. [Example](./ChessGame/Requirements.md)
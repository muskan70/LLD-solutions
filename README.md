## LLD-solutions
1. [N-elevator System](./NElevatorSystem)
2. [Build in Cache Library](./CacheSystem) : Strategy Design Pattern
3. [Gym Centre System](./GymCentreClassBookingSystem)
4. [Splitwise](./SplitwiseSystem)
5. [Parking Lot](./ParkingLot)
6. [Book Manager](./BookManager)
7. [MultiAttribute Key-Value Cache System](./MultiAttributeKeyValueCacheSystem)
8. [Course Curriculum manager](./CourseCurriculumManager)
9. [Gaming Score Board System](./GamingScoreBoardSystem)
10. [Tic Tac Toe Game](./TicTacToeGame)
11. [Movie Ticket Booking System](./MovieTicketBookingSystem)
12. [Snakes and Ladders Game](./SnakesNLaddersGame) : Factory Design Pattern
13. [Publisher Subscriber System](./PublisherSubcriberSystem) : Pull Based Architecture, Observer Design Pattern
14. [Concurrent Workers](./ConcurrencyInGo)
15. [Unix file search API](./UnixFileSearchAPI) : Composite Design Pattern, Specification Design Pattern
16. [Vending Machine](./VendingMachineSystem) : State Design Pattern, Singleton Pattern
17. [Stack Overflow](./StackOverflowSystem)
18. [Logging System](./LoggingSystem)
19. [Chat Server App](./ChatServerApp)
20. [Cricket Scoreboard](./CricketScoreboard)
21. [Cron Parser](./CronParser)
22. [Meeting Scheduler](./MeetingScheduler)

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

7. [Example](./ChessGame/Readme.md)
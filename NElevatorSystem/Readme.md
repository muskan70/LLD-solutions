## Elevator

### Requirements
1. Elevator should move up and down.
2. Types of Buttons to control elevator
- Inside Elevator to choose floor : Button Panel
- Outside Lift on Floor to call Elevator : Floor Panel
3. Dispatcher/ Scheduler

### Core Entities
1. ButtonPanel
- Attributes => isMoving, isMovingUp, startFloor, EndFloor, currentFloor
- Methods => moveToFloor(floorNo), sendInstructionToDispatcher(curFloor, destinationFloor, isMovingUp)

2. FloorPanel
- Attributes => floorNo, moveUp (boolean)
- Methods => callElevator(moveUp), sendInstructionToDistacher(floor, moveUp)

3. Elevator
- Attibutes => ElevatorId, Door, ButtonPanel, Status: MovingUp/ MovingDown, NotOperational, OnFloor, WeightLimit, Requests

4. Door
- Attributes => isOpen
- Methods => openDoor(), closeDoor()

5. Dispatcher
- Attributes => Elevators : List{Elevator}, Floors: List{FloorPanel}, NoofElevators
- Methods => startElevatorSystem(), stopElevatorSystem()
- Strategy Pattern to decide requests => 
    1. FirstComeFirstServe using queue
    2. ClosestElevatorAlgo





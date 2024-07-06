## Run this code run on terminal using following steps
Step 1: go mod init snakesNLadders
Step 2: go build
Step 3: ./snakesNLadders

###  GameDetails.json file is used to take initial inputs, use it to change any input data
- To manually enter dice values -> change "DiceEntry":"automated" to "DiceEntry":"manually" and viceversa
- 4 jumpers added : snakes, ladders, crocodiles, mines can be modified through GameDetails.json
- Multiplayer game  with no.of players  and their starting location can be modified through GameDetails.json
- MultiDice game with movement strategy (SUM, MIN, MAX) can be modified through GameDetails.json

## Factory Pattern is used to create Jumpers (snakes, ladders, crocodiles, mines)
### Classes used:
1. Player
2. Dice
3. Board
4. Game
5. Snake
6. Ladder
7. Crocodile
8. Mine
### Interface used
1. IJumper


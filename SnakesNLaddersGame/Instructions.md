# to run this code run on terminal following steps
Step 1: go mod init snakesNLadders
Step 2: go build
Step 3: ./snakesNLadders

#  GameDetails.json file is used to take initial inputs, use it to change any input data

# to manually enter dice values -> change "DiceEntry":"automated" to "DiceEntry":"manually" and viceversa

# 4 jumpers added : snakes, ladders, crocodiles, mines can be modified through GameDetails.json

# multiplayer game  with no.of players  and their starting location can be modified through GameDetails.json

# multiDice game with movement strategy (SUM, MIN, MAX) can be modified through GameDetails.json

# factory Pattern is used to create Jumpers (snakes, ladders, crocodiles, mines)

# classes used:
1. Player
2. Dice
3. Board
4. Game
5. Snake
6. Ladder
7. Crocodile
8. Mine

# interface used
1. IJumper


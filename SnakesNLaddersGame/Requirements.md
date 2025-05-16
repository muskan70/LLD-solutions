## Snakes & Ladders
**Problem Statement** : You have to design & implement a Snakes and Ladders game that supports the following functionality.

### Requirements:
- **Mandatory Requirements**: You have to take a configuration (can be a yml/json file) with the following parameters.
> Number of players: N<br>
> Board Size: BS (BS x BS)<br>
> Number of Snakes: S<br>
> Number of Ladders: L<br>
> Number of Dies: D<br>
> Movement Strategy: MS<br>
> Note: Movement strategy is either SUM (sum of numbers on dies), MAX (max of numbers on dies), MIN (min of number on dies).
- You will be given a sample input to populate the board. Post which the game has to be simulated among N players.

- **Rules**:
> - Snake always takes you to the cell where its tail is, and has to be a number less than where you are at currently.
> - Ladder takes you up (strictly).
> - If a player (A) comes to a cell where another player (B) is placed already, the previously placed player (B) has to start again from 1.

- **Optional Extensions**:
> - Using the configuration you have to generate a random valid board & devise proper rules for placing objects on the board.
> - Write unit tests to validate all implemented functionality and their edge cases.
> - Addition of special objects:
>> Crocodile, which takes you exactly 5 steps back.<br>
>> Mine which holds you for 2 turns.


- The program should take the following as input:
> Starting location of each player.<br>
> The D die values that each player rolled in a turn.

- **Guidelines**: The die roll can be implemented using a random function.

- **Sample Input and Output**:
> Input format:
> <br>● Total Snakes S
> <br>● Following S lines contains pair (Snake’s Head and Snake’s Tail)
> <br>● Total Ladders L
> <br>● Following L lines contains pair (Ladder bottom and Ladder top)
> <br>● N no of players
> <br>● Following N lines contains names & starting locations of each Player
> <br>● An override to manually enter the D die values that each player rolled in each turn.
>> 9
>> <br>62 5
>> <br>33 6
>> <br>49 9
>> <br>88 16
>> <br>41 20
>> <br>56 53
>> <br>98 64
>> <br>93 73
>> <br>95 75
>> <br>8
>> <br>2 37
>> <br>27 46
>> <br>10 32
>> <br>51 68
>> <br>61 79
>> <br>65 84
>> <br>71 91
>> <br>81 100
>> <br>2
>> <br>Gaurav 1
>> <br>Sagar 1
> Output:
>> Gaurav rolled a 6 and moved from 0 to 6
>> <br>Sagar rolled a 1 and moved from 0 to 1
>> <br>Gaurav rolled a 6 and moved from 6 to 12
>> <br>Sagar rolled a 4 and moved from 1 to 5
>> <br>Gaurav rolled a 4 and moved from 12 to 16
>> <br>Sagar rolled a 6 and moved from 5 to 11
>> <br>Gaurav rolled a 5 and moved from 16 to 21
>> <br>Sagar rolled a 4 and moved from 11 to 15
>> <br>Gaurav rolled a 1 and moved from 21 to 22
>> <br>Sagar rolled a 6 and moved from 15 to 21
>> <br>Gaurav rolled a 6 and moved from 22 to 28
>> <br>Sagar rolled a 2 and moved from 21 to 23
>> <br>Gaurav rolled a 6 and moved from 28 to 34
>> <br>Sagar rolled a 6 and moved from 23 to 29
>> <br>Gaurav rolled a 5 and moved from 34 to 39
>> <br>Sagar rolled a 2 and moved from 29 to 31
>> <br>Gaurav rolled a 2 and bitten by snake at 41 and moved from 41 to 20
>> <br>Sagar rolled a 5 and moved from 31 to 36
>> <br>Gaurav rolled a 3 and moved from 20 to 23
>> <br>Sagar rolled a 5 and bitten by snake at 41 and moved from 41 to 20
>> <br>Gaurav rolled a 6 and moved from 23 to 29
>> <br>Sagar rolled a 3 and moved from 20 to 23
>> <br>Gaurav rolled a 2 and moved from 29 to 31
>> <br>Sagar rolled a 3 and moved from 23 to 26
>> <br>Gaurav rolled a 3 and moved from 31 to 34
>> <br>Sagar rolled a 5 and moved from 26 to 31
>> <br>Gaurav rolled a 3 and moved from 34 to 37
>> <br>Sagar rolled a 4 and moved from 31 to 35
>> <br>Gaurav rolled a 2 and moved from 37 to 39
>> <br>Sagar rolled a 5 and moved from 35 to 40
>> <br>Gaurav rolled a 2 and bitten by snake at 41 and moved from 41 to 20
>> <br>Sagar rolled a 5 and moved from 40 to 45
>> <br>Gaurav rolled a 2 and moved from 20 to 22
>> <br>Sagar rolled a 6 and climbed the ladder at 51 moved from 51 to 68
>> <br>Gaurav rolled a 3 and moved from 22 to 25
>> <br>Sagar rolled a 3 and climbed the ladder at 71 and moved from 71 to 91
>> <br>Gaurav rolled a 5 and moved from 25 to 30
>> <br>Sagar rolled a 2 and bitten by snake at 93 and moved from 93 to 73
>> <br>Gaurav rolled a 5 and moved from 30 to 35
>> <br>Sagar rolled a 6 and moved from 73 to 79
>> <br>Gaurav rolled a 5 and moved from 35 to 40
>> <br>Sagar rolled a 1 and moved from 79 to 80
>> <br>Gaurav rolled a 4 and moved from 40 to 44
>> <br>Sagar rolled a 2 and moved from 80 to 82
>> <br>Gaurav rolled a 5 and bitten by snake at 49 and moved from 49 to 9
>> <br>Sagar rolled a 4 and moved from 82 to 86
>> <br>Gaurav rolled a 1 and climbed the ladder at 10 and moved from 10 to 32
>> <br>Sagar rolled a 6 and moved from 86 to 92
>> <br>Gaurav rolled a 3 and moved from 32 to 35
>> <br>Sagar rolled a 4 and moved from 92 to 96
>> <br>Gaurav rolled a 1 and moved from 35 to 36
>> <br>Sagar rolled a 1 and moved from 96 to 97
>> <br>Gaurav rolled a 1 and moved from 36 to 37
>> <br>Sagar rolled a 5 and moved from 97 to 97
>> <br>Gaurav rolled a 6 and moved from 36 to 42
>> <br>Sagar rolled a 3 and moved from 97 to 100
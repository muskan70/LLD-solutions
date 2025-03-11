### Requirements
1. System should update scores of each team after every ball, bowling as well as batting score.
2. Player score should also be updated accordingly.
3. Keep track of strike changes at the end of the over or after taking singles or 3s

### CoreComponents
1. Player
2. Team
3. Scorecard of a team
4. Scorecard of player
5. Innings
6. Over
7. Match

### Class Definations
1. Player 
- Attributes => Player Scorecard, name, status
- Methods    => UpdateBattingScore(), UpdateBowlingScore(), UpdateStatus()

2. Team
- Attributes => Name, List{Player}, Team Scorecard, Status

3. Innings
- Attributes => BattingTeam, BattingScores, BowlingScores, OverDetails, CurrentStrike, CurrentBattingPlayers[2]
- Methods    => AddOver(), StartInnings() 

4. Match
- Attributes => TeamA, TeamB, Innings[2], TossStatus, TossAction, Winner
- Attributes => PlayMatch()

5. Scorecard
- Attributes => Runs, BallsPlayed, Sixes, Fours

6. Over
- Attributes => []Balls


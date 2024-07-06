package game

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"snakesNLadders/board"
	"snakesNLadders/constants"
	"snakesNLadders/dice"
	"snakesNLadders/jumperFactory"
	"snakesNLadders/player"
)

type InputPlayer struct {
	Name             string
	StartingLocation int
}

type GameInput struct {
	BoardSize          int
	NumberOfSnakes     int
	SnakesPosition     []jumperFactory.Jumper
	NumberOfLadders    int
	LaddersPosition    []jumperFactory.Jumper
	NumberOfCrocodiles int
	CrocodilesPosition []int
	NumberOfMines      int
	MinesPosition      []int
	NumberOfDies       int
	MovementStrategy   string
	NumberOfPlayers    int
	PlayersDetails     []InputPlayer
	DiceEntry          string
}

func InitiateGame() *Game {
	jsonFile, err1 := os.Open("GameDetails.json")
	if err1 != nil {
		log.Println(err1)
		return nil
	}
	defer jsonFile.Close()

	byteValue, err2 := io.ReadAll(jsonFile)
	if err2 != nil {
		log.Println("Invalid json file format", err2)
		return nil
	}

	var gameInputs GameInput
	err := json.Unmarshal(byteValue, &gameInputs)
	if err != nil {
		log.Println("Invalid json file format", err)
		return nil
	}

	var players player.Players
	for i := 0; i < gameInputs.NumberOfPlayers; i++ {
		players = append(players, player.NewPlayer(gameInputs.PlayersDetails[i].Name, gameInputs.PlayersDetails[i].StartingLocation))
	}

	gameBoard := board.NewBoard(gameInputs.BoardSize)
	if err = gameBoard.AddSnakes(gameInputs.NumberOfSnakes, gameInputs.SnakesPosition); err != nil {
		log.Println("Invalid game inputs : ", err)
		return nil
	}
	if err = gameBoard.AddLadders(gameInputs.NumberOfLadders, gameInputs.LaddersPosition); err != nil {
		log.Println("Invalid game inputs : ", err)
		return nil
	}
	if err = gameBoard.AddCrocodiles(gameInputs.NumberOfCrocodiles, gameInputs.CrocodilesPosition); err != nil {
		log.Println("Invalid game inputs : ", err)
		return nil
	}
	if err = gameBoard.AddMines(gameInputs.NumberOfMines, gameInputs.MinesPosition); err != nil {
		log.Println("Invalid game inputs : ", err)
		return nil
	}

	dice := dice.NewDice(gameInputs.NumberOfDies, gameInputs.MovementStrategy, gameInputs.DiceEntry)
	snakeNLaddersGame := &Game{Players: players, Board: gameBoard, Dice: dice, Status: constants.GAME_STATUS_NOT_STARTED}
	log.Println("Game successfully Initialized")
	return snakeNLaddersGame
}

# Problem Statement: Design a Chess Game where two players can play against each other.

## Step-by-Step Approach

### Understand the Problem (5 minutes) -> Clarify Requirements:
- How are moves made (input method)?
- Are there any special rules to consider (like castling, en passant)?
- Should the game handle check, checkmate, and stalemate conditions?
- Is there a need for a graphical interface or is it purely text-based?

### Identify Core Components (10 minutes)
- Player: Represents each player in the game.
- ChessBoard: Represents the 8x8 grid.
- Piece: Represents each chess piece (e.g., King, Queen, Rook).
- Move: Represents a move in the game.
- Game: Manages the overall game logic.

### Define Class Diagrams (15 minutes)
1. Player:
- Attributes => name, color
- Methods    => makeMove()

2. ChessBoard:
- Attributes => board[8][8], Piece: King(2), Queen(2), Rook(2), Bishop(2), Knight(2), Pawn(16) -> half-half for each player
- Methods    => setupBoard(), movePiece(), isCheck(), isCheckmate()

3. Piece:
- Attributes => position, color
- Methods    => validMoves()
- Subclasses => King, Queen, Rook, Bishop, Knight, Pawn

4. Move:
- Attributes => startPosition, endPosition, piece
- Methods    => execute()

5. Game:
- Attributes => players, currentTurn, chessBoard
- Methods    => startGame(), endGame(), switchTurn()

### Design Patterns (10 minutes)
- Singleton: Ensure only one instance of the Game.
- Factory: To create different pieces (King, Queen, etc.).
- Observer: To notify players about game state changes.

### Implementation Plan (5 minutes)
- Setup Classes: Create basic class structures.
- Define Methods: Implement methods for moving pieces, validating moves, and checking game states.
- Time Management:
- Class design: 15 minutes
- Method implementation: 20 minutes
- Testing and final touch-ups: 5 minutes

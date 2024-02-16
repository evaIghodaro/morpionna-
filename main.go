package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	boardSize = 9
	numRows   = 3
	numCols   = 3
)

func main() {
	// Create the tic-tac-toe board
	board := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	// Print the initial board
	printBoard(board)

	// Play the game
	play(&board)
}

func printBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" ", board[i][j], " ")
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func play(board *[3][3]string) {
	scanner := bufio.NewScanner(os.Stdin)
	player := "O"

	for {
		fmt.Printf("Player %s, enter the cell number (from 1 to 9): ", player)
		scanner.Scan()
		input := scanner.Text()

		// Check if the input is a valid number
		n, err := strconv.Atoi(input)
		if err != nil || n < 1 || n > 9 {
			fmt.Println("Enter a valid number between 1 and 9")
			continue
		}

		// Convert the cell number to board coordinates
		row := (n - 1) / 3
		col := (n - 1) % 3

		// Check if the cell is already occupied
		if (*board)[row][col] != " " {
			fmt.Println("Cell already occupied, please choose another cell.")
			continue
		}

		// Place the player's symbol in the corresponding cell of the board
		(*board)[row][col] = player

		// Print the updated board
		printBoard(*board)

		// Check if there's a winner or if the game is a draw
		if win(player == "O", convertBoardToArray(*board), "O", "X") {
			fmt.Printf("Player %s wins!\n", player)
			break
		} else if isDraw(*board) {
			fmt.Println("It's a draw!")
			break
		}

		// Switch player for the next turn
		if player == "O" {
			player = "X"
		} else {
			player = "O"
		}
	}
}

func win(player1 bool, ticTacToeBoard [boardSize]string, playerSymbol1, playerSymbol2 string) bool {
	// Determine the symbol for the current player
	playerSymbol := playerSymbol1
	if !player1 {
		playerSymbol = playerSymbol2
	}

	// Check rows
	for i := 0; i < numRows; i++ {
		if ticTacToeBoard[i*numCols] == playerSymbol && ticTacToeBoard[i*numCols+1] == playerSymbol && ticTacToeBoard[i*numCols+2] == playerSymbol {
			return true // If a row is filled with the player's symbol, return true
		}
	}

	// Check columns
	for i := 0; i < numCols; i++ {
		if ticTacToeBoard[i] == playerSymbol && ticTacToeBoard[i+3] == playerSymbol && ticTacToeBoard[i+6] == playerSymbol {
			return true // If a column is filled with the player's symbol, return true
		}
	}

	// Check diagonals
	if ticTacToeBoard[0] == playerSymbol && ticTacToeBoard[4] == playerSymbol && ticTacToeBoard[8] == playerSymbol {
		return true // If the main diagonal is filled with the player's symbol, return true
	}
	if ticTacToeBoard[2] == playerSymbol && ticTacToeBoard[4] == playerSymbol && ticTacToeBoard[6] == playerSymbol {
		return true // If the secondary diagonal is filled with the player's symbol, return true
	}

	return false // If no winning condition is met, return false
}

func convertBoardToArray(board [3][3]string) [boardSize]string {
	var array [boardSize]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			array[i*3+j] = board[i][j]
		}
	}
	return array
}

func isDraw(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create the board of the morpion
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
		fmt.Printf("Joueur %s, entrez le numéro de la case (de 1 à 9) : ", player)
		scanner.Scan()
		input := scanner.Text()

		// Vérifier si l'entrée est un nombre valide
		n, err := strconv.Atoi(input)
		if err != nil || n < 1 || n > 9 {
			fmt.Println("Entrez un nombre valide entre 1 et 9")
			continue
		}

		// Convertir le numéro de la case en coordonnées du tableau
		row := (n - 1) / 3
		col := (n - 1) % 3

		// Vérifier si la case est déjà occupée
		if (*board)[row][col] != " " {
			fmt.Println("Case déjà occupée, veuillez choisir une autre case.")
			continue
		}

		// Placer le symbole du joueur dans la case correspondante du tableau
		(*board)[row][col] = player

		// Afficher le tableau mis à jour
		printBoard(*board)

		// Vérifier s'il y a un gagnant ou si le match est nul
		// (code à ajouter ici)

		// Changer de joueur pour le prochain tour
		if player == "O" {
			player = "X"
		} else {
			player = "O"
		}
	}
}

//faire un scanne pour demander a l'utilisateur ce qu'il veut mettre en chiffre pour le placer a la case du chiffre et convertir le chiffre par 0

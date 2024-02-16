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
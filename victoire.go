const (
	tailleDamier = 9 // Taille du damier du morpion (3x3)
	nbLignes     = 3 // Nombre de lignes dans le morpion
	nbColonnes   = 3 // Nombre de colonnes dans le morpion
)

func gagner(joueur1 bool, tableauMorpion [tailleDamier]string, symboleJoueur1, symboleJoueur2 string) bool {
	symboleJoueur := symboleJoueur1
	if !joueur1 {
		symboleJoueur = symboleJoueur2
	}

	// Vérifier les lignes
	for i := 0; i < nbLignes; i++ {
		if tableauMorpion[i*nbColonnes] == symboleJoueur && tableauMorpion[i*nbColonnes+1] == symboleJoueur && tableauMorpion[i*nbColonnes+2] == symboleJoueur {
			return true
		}
	}

	// Vérifier les colonnes
	for i := 0; i < nbColonnes; i++ {
		if tableauMorpion[i] == symboleJoueur && tableauMorpion[i+3] == symboleJoueur && tableauMorpion[i+6] == symboleJoueur {
			return true
		}
	}

	// Vérifier les diagonales
	if tableauMorpion[0] == symboleJoueur && tableauMorpion[4] == symboleJoueur && tableauMorpion[8] == symboleJoueur {
		return true
	}
	if tableauMorpion[2] == symboleJoueur && tableauMorpion[4] == symboleJoueur && tableauMorpion[6] == symboleJoueur {
		return true
	}

	return false
}

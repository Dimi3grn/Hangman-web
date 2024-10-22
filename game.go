package hangman

import (
	"fmt"
)

func PlayGame(hiddenWord string, display []rune) {
	isRunning := true
	tries := 6                     // Nombre limité d'essais
	attemptedLetters := []string{} // Liste des lettres déjà tentées
	attemptedWords := []string{}
	fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
	fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage initial

	for isRunning && tries > 0 {
		printHangman(tries)
		var option string
		fmt.Println("Choisissez une lettre ou proposez un mot entier (tentatives restantes :", tries, ")")
		fmt.Scan(&option)

		// Vérifier si l'utilisateur a proposé un mot entier
		if len(option) > 1 {
			if contains(attemptedWords, option) {
				Clear()
				fmt.Println("Vous avez déjà proposé ce mot, essayez-en un autre.")
				fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
				fmt.Println("Affichage actuel :", string(display))
				continue
			}
			if option == hiddenWord {
				// Si le mot est correct, on met à jour l'affichage et on termine le jeu
				for i, char := range hiddenWord {
					display[i] = char
				}
				isRunning = false
				fmt.Println("Vous avez deviné le mot, bien joué à vous !")
				break
			} else {
				// Si le mot est incorrect, retire deux essais
				Clear()
				tries -= 2
				attemptedWords = append(attemptedWords, option)
				if tries <= 0 {
					fmt.Println("Dommage ! Vous avez épuisé vos tentatives. Le mot était :", hiddenWord)
					break
				}
				fmt.Println("Mauvaise proposition ! Deux tentatives en moins.")
				fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
				fmt.Println("Affichage actuel :", string(display))
				continue
			}
		}

		// Vérifier si la lettre a déjà été tentée
		if contains(attemptedLetters, option) {
			Clear()
			fmt.Println("Vous avez déjà choisi cette lettre, essayez-en une autre.")
			fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
			fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage initial
			continue                                           // Ne pas réduire le nombre d'essais
		}

		// Ajouter la lettre à la liste des lettres tentées
		attemptedLetters = append(attemptedLetters, option)

		// Mettre à jour l'affichage
		if !UpdateDisplay(hiddenWord, display, option) {
			Clear()
			tries-- // Décrémente le nombre d'essais si la lettre n'est pas trouvée
			fmt.Println("Mauvaise lettre !")
		}

		Clear()
		fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage mis à jour
		// Vérifier les conditions de victoire uniquement ici
		if CheckComp(display) == false {
			isRunning = false
			fmt.Println("Vous avez deviné le mot, bien joué à vous !")
			// On ne vérifie pas ici, car le CheckComp doit s'occuper de l'affichage
		} else if tries == 0 {
			fmt.Println("Dommage ! Vous avez épuisé vos tentatives. Le mot était :", hiddenWord)
		}
	}

	// Vérifier si le joueur a gagné après la dernière mise à jour
}

func contains(attemptedLetters []string, letter string) bool {
	for _, attempted := range attemptedLetters {
		if attempted == letter {
			return true
		}
	}
	return false
}

func UpdateDisplay(hiddenWord string, display []rune, option string) bool {
	correctGuess := false
	for k, char := range hiddenWord {
		if option == string(char) {
			display[k] = char
			correctGuess = true
		}
	}
	return correctGuess // Renvoie true si la lettre est correcte
}

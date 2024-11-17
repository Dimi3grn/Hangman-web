package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpdatePoints(user string, pointsToAdd int) error {
	// Lire le contenu du fichier
	file, err := os.Open("./view/assets/logins/loggins.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Parcourir chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		// Diviser chaque ligne en composants {user} {password} {points}
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}

		username := parts[0]
		password := parts[1]
		points, err := strconv.Atoi(parts[2])
		if err != nil {
			return err
		}

		// Si l'utilisateur correspond, ajouter les points
		if username == user {
			points += pointsToAdd
		}

		// Reconstituer la ligne et l'ajouter à la liste
		updatedLine := fmt.Sprintf("%s %s %d", username, password, points)
		lines = append(lines, updatedLine)
	}

	// Écrire le contenu mis à jour dans le fichier
	file.Close()
	outputFile, err := os.Create("./view/assets/logins/loggins.txt")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range lines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

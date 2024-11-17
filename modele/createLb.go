package hangman

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type User struct {
	Name     string
	Password string
	Points   int
}

func readUsersFromFile(filename string) ([]User, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue // Ignorer les lignes mal formatées
		}

		// Conversion du score en entier
		points, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		// Ajouter l'utilisateur à la liste
		user := User{Name: parts[0], Password: parts[1], Points: points}
		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func writeLeaderboardToFile(filename string, users []User) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, user := range users {
		line := fmt.Sprintf("%s %d\n", user.Name, user.Points)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// Fonction principale pour générer le leaderboard
func GenerateLeaderboard(inputFile, outputFile string) error {
	// Lire les utilisateurs depuis le fichier
	users, err := readUsersFromFile(inputFile)
	if err != nil {
		return err
	}

	// Trier les utilisateurs par points (ordre décroissant)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})

	// Prendre les 7 meilleurs utilisateurs
	topUsers := users
	if len(users) > 7 {
		topUsers = users[:7]
	}

	// Écrire le leaderboard dans le fichier de sortie
	err = writeLeaderboardToFile(outputFile, topUsers)
	if err != nil {
		return err
	}

	fmt.Println("Leaderboard généré avec succès !")
	return nil
}

type Position struct {
	Name  string
	Score int
}

// Structure pour le leaderboard avec 7 positions
type Leaderboard struct {
	Pos1 Position
	Pos2 Position
	Pos3 Position
	Pos4 Position
	Pos5 Position
	Pos6 Position
	Pos7 Position
}

// Fonction pour lire le fichier leaderboard.txt et remplir la structure Leaderboard
func ReadLeaderboard(filename string) (Leaderboard, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Leaderboard{}, err
	}
	defer file.Close()

	var leaderboard Leaderboard
	positions := []*Position{&leaderboard.Pos1, &leaderboard.Pos2, &leaderboard.Pos3, &leaderboard.Pos4, &leaderboard.Pos5, &leaderboard.Pos6, &leaderboard.Pos7}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < 7 {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		score, err := strconv.Atoi(parts[1])
		if err != nil {
			return Leaderboard{}, err
		}

		positions[i].Name = name
		positions[i].Score = score
		i++
	}

	if err := scanner.Err(); err != nil {
		return Leaderboard{}, err
	}

	return leaderboard, nil
}

package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func userExists(username string) (bool, error) {
	// Ouvrir le fichier pour lecture
	file, err := os.Open("./view/assets/logins/loggins.txt")
	if err != nil {
		return false, fmt.Errorf("erreur d'ouverture du fichier: %v", err)
	}
	defer file.Close()

	// Lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Vérifier si le nom d'utilisateur existe déjà
		parts := strings.Split(line, " ")
		if len(parts) == 2 && parts[0] == username {
			return true, nil // Utilisateur trouvé
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("erreur de lecture du fichier: %v", err)
	}

	return false, nil // Utilisateur non trouvé
}

// Fonction qui écrit un nom d'utilisateur et un mot de passe dans le fichier
func WriteCredentials(username, password string) error {

	exists, err := userExists(username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("l'utilisateur '%s' existe déjà dans le fichier", username)
	}

	// Ouvrir ou créer le fichier pour ajouter du contenu
	file, err := os.OpenFile("./view/assets/logins/loggins.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erreur d'ouverture du fichier: %v", err)
	}
	defer file.Close()

	// Écrire l'utilisateur et le mot de passe dans le fichier
	_, err = fmt.Fprintf(file, "%s %s\n", username, password)
	if err != nil {
		return fmt.Errorf("erreur d'écriture dans le fichier: %v", err)
	}

	fmt.Println("Utilisateur et mot de passe ajoutés avec succès.")
	return nil
}

package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to verify if username and password match any entry in the file
func VerifyCredentials(username, password string) bool {
	// Open the file
	fmt.Println(os.Getwd())
	file, err := os.Open("./view/assets/logins/loggins.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into username and password
		parts := strings.Split(line, " ")
		if len(parts) == 2 && parts[0] == username && parts[1] == password {
			return true
		}
	}

	// If no match was found
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return false
}

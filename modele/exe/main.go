package main

//si tu veux ajouter un nouveau fichier avec des mots, créé un fichier .txt, ouvre le et met a la première ligne la difficultée de ton fichier de mots
//ensuite copie colle tes mots, chaque mot est différentié de l'autre pas un saut a la ligne, pas de majuscules, d'accents ou autres caractères spéciaux
//  --- pas d'espaces j'ai essayer mais flm c trop chiant avec bufflo reader donc bon

import (
	// replace with your actual module name
	hangman "Hangman/modele"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {

	temp, err := template.ParseGlob("./view/template/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("erreur %s", err.Error()))
		return
	}

	http.HandleFunc("/landingPage", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Landing", nil)
	})

	http.HandleFunc("/hangman/mainGame", func(w http.ResponseWriter, r *http.Request) {
		type HangmanPage struct {
			MotCache string
			Display  string
		}
		displayData := HangmanPage{}

		fileName := "halloween.txt"
		wordsArr := hangman.ReadWordsFromFile(fileName)

		hiddenWord := hangman.SelectRandomWord(wordsArr)

		displayData.MotCache = hiddenWord
		displayData.Display = string(hangman.InitializeDisplay(hiddenWord))
		fmt.Fprintln(os.Stdout, displayData.MotCache)
		temp.ExecuteTemplate(w, "Hangman", displayData)
	})

	http.HandleFunc("/hangman/treatment", func(w http.ResponseWriter, r *http.Request) {

	})

	fileServer := http.FileServer(http.Dir("./view/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8080", nil)
}

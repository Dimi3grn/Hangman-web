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

	type HangmanPage struct {
		MotCache   string
		Display    string
		Tries      int
		AttWrods   []string
		AttLetters []string
	}
	displayData := HangmanPage{}

	http.HandleFunc("/landingPage", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Landing", nil)
	})

	http.HandleFunc("/landingPage/treatment", func(w http.ResponseWriter, r *http.Request) {
		fileName := "halloween.txt"
		wordsArr := hangman.ReadWordsFromFile(fileName)
		fmt.Fprintln(os.Stdout, wordsArr)
		hiddenWord := hangman.SelectRandomWord(wordsArr)
		fmt.Fprintln(os.Stdout, hiddenWord)

		fmt.Fprintln(os.Stdout, hangman.InitializeDisplay(hiddenWord))

		displayData = HangmanPage{hiddenWord, string(hangman.InitializeDisplay((hiddenWord))), 6, []string{}, []string{}}
		fmt.Fprintln(os.Stdout, hiddenWord)

		http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)

	})

	http.HandleFunc("/hangman/mainGame", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(os.Stdout, displayData.MotCache)
		temp.ExecuteTemplate(w, "Hangman", displayData)
	})

	http.HandleFunc("/hangman/treatment", func(w http.ResponseWriter, r *http.Request) {

		hangman.Verify(r.FormValue("mot"), displayData.AttWrods, displayData.AttLetters, displayData.MotCache, displayData.Display, displayData.Tries)
		fmt.Fprintln(os.Stdout, displayData.Display)
		fmt.Fprintln(os.Stdout, displayData.Tries)
		http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)
	})

	fileServer := http.FileServer(http.Dir("./view/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8080", nil)
}

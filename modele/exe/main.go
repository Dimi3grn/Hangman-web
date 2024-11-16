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
		MotCache       string
		Display        string
		Tries          int
		AttWrods       []string
		AttLetters     []string
		ImgPath        string
		IsSolved       bool
		DisplayMessage string
	}

	type LogPage struct {
		PageMessage string
		LoggedIn    bool
	}

	PageData := LogPage{}
	displayData := HangmanPage{}

	PageData = LogPage{PageMessage: "", LoggedIn: false}

	//Page D'accueil
	http.HandleFunc("/landingPage", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, "Landing Page")

		temp.ExecuteTemplate(w, "landing", PageData)
	})

	http.HandleFunc("/signup/treatment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, "Login Treatment")
		login := r.FormValue("login")
		password := r.FormValue("password")
		fmt.Fprintln(os.Stdout, login)
		fmt.Fprintln(os.Stdout, password)
		hangman.WriteCredentials(login, password)
		http.Redirect(w, r, "/landingPage", http.StatusSeeOther)
	})

	http.HandleFunc("/login/treatment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, "Login Treatment")
		login := r.FormValue("login")
		password := r.FormValue("password")
		fmt.Fprintln(os.Stdout, login)
		fmt.Fprintln(os.Stdout, password)
		if hangman.VerifyCredentials(login, password) {
			PageData.LoggedIn = true
			PageData.PageMessage = "Login Success"
			fmt.Fprintln(os.Stdout, "Login Success")
		} else {
			PageData.LoggedIn = false
			PageData.PageMessage = "Login Failed"
			fmt.Fprintln(os.Stdout, "Login Failed")
		}

		http.Redirect(w, r, "/landingPage", http.StatusSeeOther)
	})

	//Choisir le thème du mot
	http.HandleFunc("/landingPage/treatment", func(w http.ResponseWriter, r *http.Request) {
		if !PageData.LoggedIn {
			PageData.PageMessage = "Login required to play"
			http.Redirect(w, r, "/landingPage", http.StatusSeeOther)
		} else {
			fmt.Fprintln(os.Stdout, "Landing Page Treatment")

			fileName := r.FormValue("formID") + ".txt"
			wordsArr := hangman.ReadWordsFromFile(fileName)
			fmt.Fprintln(os.Stdout, wordsArr)
			hiddenWord := hangman.SelectRandomWord(wordsArr)
			fmt.Fprintln(os.Stdout, hiddenWord)

			fmt.Fprintln(os.Stdout, hangman.InitializeDisplay(hiddenWord))

			displayData = HangmanPage{hiddenWord, string(hangman.InitializeDisplay((hiddenWord))), 6, []string{}, []string{}, "/static/img/r6-operators-list-" + hiddenWord + ".avif", false, ""}
			fmt.Fprintln(os.Stdout, hiddenWord)
			fmt.Fprintln(os.Stdout, displayData.ImgPath)

			http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)
		}

	})

	// Page de Jeux
	http.HandleFunc("/hangman/mainGame", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		fmt.Fprintln(os.Stdout, "Main Game")

		fmt.Fprintln(os.Stdout, displayData.MotCache)
		temp.ExecuteTemplate(w, "Hangman", displayData)
	})

	http.HandleFunc("/hangman/treatment", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		fmt.Fprintln(os.Stdout, "Treatment")

		displayData.Display, displayData.Tries, displayData.IsSolved, displayData.DisplayMessage = hangman.Verify(r.FormValue("mot"), &displayData.AttWrods, &displayData.AttLetters, displayData.MotCache, displayData.Display, displayData.Tries, displayData.IsSolved)
		fmt.Fprintln(os.Stdout, displayData.Display)
		fmt.Fprintln(os.Stdout, displayData.Tries)
		fmt.Fprintln(os.Stdout, displayData.IsSolved)
		http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)
	})

	fmt.Fprintln(os.Stdout, "Serveur démarré sur http://localhost:8085")
	fileServer := http.FileServer(http.Dir("./view/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8085", nil)

}

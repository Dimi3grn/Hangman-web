package main

import (
	hangman "Hangman/modele"
	"log"
	"net/http"
	"text/template"
)

type HangmanPage struct {
	MotCache   string
	Display    string
	Tries      int
	AttWrods   []string
	AttLetters []string
}

var (
	displayData HangmanPage
	temp        *template.Template
)

func main() {
	var err error
	temp, err = template.ParseGlob("./../../view/template/*.html")
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates : %v", err)
	}

	http.HandleFunc("/accueil", accueilHandler)
	http.HandleFunc("/landingPage/treatment/halloween", func(w http.ResponseWriter, r *http.Request) {
		handleThemeSelection(w, r, "halloween.txt")
	})
	http.HandleFunc("/landingPage/treatment/noel", func(w http.ResponseWriter, r *http.Request) {
		handleThemeSelection(w, r, "noel.txt")
	})
	http.HandleFunc("/landingPage/treatment/summer", func(w http.ResponseWriter, r *http.Request) {
		handleThemeSelection(w, r, "summer.txt")
	})
	http.HandleFunc("/hangman/mainGame", mainGameHandler)
	http.HandleFunc("/hangman/treatment", treatmentHandler)

	// Servir les fichiers statiques
	fileServer := http.FileServer(http.Dir("./../../view/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Page d'accueil
func accueilHandler(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "accueil", nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page d'accueil", http.StatusInternalServerError)
	}
}

// Gestion de la sélection de thème
func handleThemeSelection(w http.ResponseWriter, r *http.Request, fileName string) {
	wordsArr := hangman.ReadWordsFromFile(fileName)
	if wordsArr == nil {
		log.Printf("Erreur lors de la lecture du fichier %s", fileName)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	hiddenWord := hangman.SelectRandomWord(wordsArr)
	display := hangman.InitializeDisplay(hiddenWord)

	displayData = HangmanPage{
		MotCache:   hiddenWord,
		Display:    string(display),
		Tries:      6,
		AttWrods:   []string{},
		AttLetters: []string{},
	}

	log.Printf("Mot sélectionné : %s", hiddenWord)
	http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)
}

// Page de jeu
func mainGameHandler(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "jeu", displayData)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page de jeu", http.StatusInternalServerError)
	}
}

// Traitement du jeu
func treatmentHandler(w http.ResponseWriter, r *http.Request) {
	mot := r.FormValue("mot")
	displayData.Display, displayData.Tries = hangman.Verify(
		mot,
		&displayData.AttWrods,
		&displayData.AttLetters,
		displayData.MotCache,
		displayData.Display,
		displayData.Tries,
	)

	log.Printf("Mot affiché : %s, Essais restants : %d", displayData.Display, displayData.Tries)
	http.Redirect(w, r, "/hangman/mainGame", http.StatusSeeOther)
}

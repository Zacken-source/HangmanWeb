package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataForm struct {
	Motsecret       string
	Motcachee       string
	LettresUtilisee string
	Victoire        bool
	Essaies         int
	Echec           bool
}

func (s *Structure) web() {
	// chargement de tous les répertoirs présents dans "Hangman-Web"
	http.Handle("/hangmanstage/", http.StripPrefix("/hangmanstage/", http.FileServer(http.Dir("hangmanstage"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))
	http.Handle("/texte/", http.StripPrefix("/texte/", http.FileServer(http.Dir("texte"))))
	http.Handle("/HtmlCss/", http.StripPrefix("/HtmlCss/", http.FileServer(http.Dir("HtmlCss"))))

	http.HandleFunc("/", s.home)
	http.HandleFunc("/home.html", s.home)
	http.HandleFunc("/game.html", s.hangman)
	//http.HandleFunc("/secondgame", s.sgame)
	// chargement du port utilisé
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}

// fonctions pour chaque page

func (s *Structure) home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("HtmlCss/home.html"))
	tmpl.Execute(w, nil)
}

func (s *Structure) hangman(w http.ResponseWriter, r *http.Request) {
	s.TheGame(w, r)
}

func (s *Structure) TheGame(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("HtmlCss/game.html"))
	letter := r.Form.Get("letter")
	s.Letter = []rune(letter)
	var check bool = false

	if len(s.Letter) == 1 && s.CheckLetter(s.Letter) {
		check = true
	} else if len(s.Letter) > 1 && s.CheckWord(s.Letter) {
		check = true
	}

	if !check {
		s.Lives -= 1
	}

	s.CheckOut()

	web := DataForm{
		Motsecret:       string(s.SecretWord),
		Motcachee:       string(s.Blanks),
		LettresUtilisee: s.LetterTested,
		Victoire:        s.Win,
		Essaies:         s.Lives,
		Echec:           s.Lose,
	}
	tmpl.Execute(w, web)
}

package hangman

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// verifie si la lettre proposée l'a déjà été et rajout dans lettertested (renvoie true ou false)
func (s *Structure) VerifLetter(letter []rune) bool {
	isHere := true
	for _, i := range s.LetterTested {
		if i == letter[0] {
			isHere = false
		}
	}
	if isHere {
		s.LetterTested += "," + string(letter)
	}
	return isHere
}

// verifie si la lettre est dans le mot et le remplace dans le bon cas (renvoie true ou false)
func (s *Structure) CheckLetter(letter []rune) bool {
	isHere := false
	for _, l := range letter {
		for index, r := range s.SecretWord {
			if l == r {
				s.Blanks[index] = r
				isHere = true
			}
		}
	}
	return isHere
}

// verifie si le mot est le bon
func (s *Structure) CheckWord(letter []rune) bool {
	var count1 int = len(s.SecretWord)
	var count2 int = 0
	for index, r := range s.SecretWord {
		if letter[index] == r {
			count2 += 1
		}
	}
	isHere := false
	if count1 == count2 {
		isHere = true
	}
	return isHere
}

// Sélectionner un mot au hasard
func (s *Structure) SelectRandomWord() []rune {
	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile("hangmanClassic/texte/dictionnaire.txt")
	if err != nil {
		return nil
	}
	ListOfWord := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(ListOfWord))
	word := ListOfWord[randomIndex]
	return []rune(word)
}

// Initialiser les blancs
func (s *Structure) InitializeBlanks() []rune {
	blanks := make([]rune, len(s.SecretWord))
	for i := range blanks {
		blanks[i] = '_'
	}
	return blanks
}

// verifie le nb de vie / verifie si le mot est trouvé
func (s *Structure) CheckOut() {
	if s.Lives == 0 {
		s.Running = false
		s.Lose = true
	}
	var count = len(s.SecretWord)
	var count2 int = 0
	for _, i := range s.SecretWord {
		if s.Blanks[i] == rune(i) {
			count2 += 1
		}
	}
	if count2 == count {
		s.Win = true
	}
}

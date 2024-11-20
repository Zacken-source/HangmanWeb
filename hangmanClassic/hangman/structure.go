package hangman

type Structure struct {
	//variable permettant de lancer le jeu
	Running bool

	//variables permettant la résolution du hangman
	Letter       []rune
	SecretWord   []rune
	Blanks       []rune
	Lives        int
	LetterTested string

	//variables servant de changer l'état du jeu
	Win  bool
	Lose bool
}

func (s *Structure) Run() {
	s.init()
	if s.Running {
		s.web()
	}
}

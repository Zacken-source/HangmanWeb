package hangman

func (s *Structure) init() {

	s.Running = true

	s.Letter = []rune{}
	s.SecretWord = s.SelectRandomWord()
	s.Blanks = s.InitializeBlanks()
	s.Lives = 10
	s.LetterTested = ""

	s.Win = false
	s.Lose = false

}

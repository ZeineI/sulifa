package models

var GameChoices = [3]string{"r", "p", "s"}

var GameChoiceNameMap = map[string]string{
	"r": "Rock",
	"s": "Scissors",
	"p": "Paper",
}

var GameWinRules = map[string]string{
	"r": "s",
	"p": "r",
	"s": "p",
}

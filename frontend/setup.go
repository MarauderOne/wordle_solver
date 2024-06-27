package main

import (
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"strings"
)

// Define a data structure for user inputs
type BoxData struct {
	Character string `json:"character"`
	Color     string `json:"color"`
}

// Define a function to initialise the answerList
func createNewAnswerList() (dictionary *dictionary_tools.MySimpleDict) {
	d := dictionary_tools.NewSimpleDict()
	d.Load("../dictionary_tools/initialList.dict")
	return d
}

// Define a function to revise the answerList based on given regex patterns
func reviseAnswerList(answersList *dictionary_tools.MySimpleDict, regexPattern string) (revisedDictionary *dictionary_tools.MySimpleDict) {
	newAnswerList := answersList.Lookup(regexPattern, 0, 6000)
	d := dictionary_tools.NewSimpleDict()
	d.AddWordsList(newAnswerList)
	return d
}

// Define a function to determine if a character is not alphabetic
func nonAlpha(char string) bool {

	//Define a list of alphabetic characters
	const alpha = "abcdefghijklmnopqrstuvwxyz"

	if (strings.Contains(alpha, strings.ToLower(string(char)))) || (string(char) == "") {
		return false
	} else {
		return true
	}
}

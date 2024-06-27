package main

import (
	"fmt"
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

//Define a function to set regex lookup patterns
func setRegexPatterns(i int, char string) (greenRegex, yellowRegex, greyRegex string) {
	//Assign position based on index number in the array (i)
	switch i {
	case 0, 5, 10, 15, 20, 25:
		//Character in first position
		greenRegex = fmt.Sprintf("%v....", char)
		yellowRegex = fmt.Sprintf("[^%v][%v{1,}]...$|^[^%v].[%v{1,}]..$|^[^%v]..[%v{1,}].$|^[^%v]...[%v{1,}]", char, char, char, char, char, char, char, char)
		greyRegex = fmt.Sprintf("[^%v]....", char)
	case 1, 6, 11, 16, 21, 26:
		//Character in second position
		greenRegex = fmt.Sprintf(".%v...", char)
		yellowRegex = fmt.Sprintf("[%v{1,}][^%v]...$|^.[^%v][%v{1,}]..$|^.[^%v].[%v{1,}].$|^.[^%v]..[%v{1,}]", char, char, char, char, char, char, char, char)
		greyRegex = fmt.Sprintf(".[^%v]...", char)
	case 2, 7, 12, 17, 22, 27:
		//Character in third position
		greenRegex = fmt.Sprintf("..%v..", char)
		yellowRegex = fmt.Sprintf("[%v{1,}].[^%v]..$|^.[%v{1,}][^%v]..$|^..[^%v][%v{1,}].$|^..[^%v].[%v{1,}]", char, char, char, char, char, char, char, char)
		greyRegex = fmt.Sprintf("..[^%v]..", char)
	case 3, 8, 13, 18, 23, 28:
		//Character in fourth position
		greenRegex = fmt.Sprintf("...%v.", char)
		yellowRegex = fmt.Sprintf("[%v{1,}]..[^%v].$|^.[%v{1,}].[^%v].$|^..[%v{1,}][^%v].$|^...[^%v][%v{1,}]", char, char, char, char, char, char, char, char)
		greyRegex = fmt.Sprintf("...[^%v].", char)
	case 4, 9, 14, 19, 24, 29:
		//Character in fifth position
		greenRegex = fmt.Sprintf("....%v", char)
		yellowRegex = fmt.Sprintf("[%v{1,}]...[^%v]$|^.[%v{1,}]..[^%v]$|^..[%v{1,}].[^%v]$|^...[%v{1,}][^%v]", char, char, char, char, char, char, char, char)
		greyRegex = fmt.Sprintf("....[^%v]", char)
	}
	return greenRegex, yellowRegex, greyRegex
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

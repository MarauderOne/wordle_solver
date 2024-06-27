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
func setRegexPatterns(i int, char string, gridData []BoxData) (greenRegex, yellowRegex, greyRegex string) {
	//Assign position based on index number in the array (i)
	switch i {
	case 0, 5, 10, 15, 20, 25:
		//Character in first position
		greenRegex = fmt.Sprintf("%v....", char)
		yellowRegex = fmt.Sprintf("[^%v][%v{1,}]...$|^[^%v].[%v{1,}]..$|^[^%v]..[%v{1,}].$|^[^%v]...[%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i+1].Character != char) && (gridData[i+2].Character != char) && (gridData[i+3].Character != char) && (gridData[i+4].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			greyRegex = fmt.Sprintf("[^%v]....", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 1, 6, 11, 16, 21, 26:
		//Character in second position
		greenRegex = fmt.Sprintf(".%v...", char)
		yellowRegex = fmt.Sprintf("[%v{1,}][^%v]...$|^.[^%v][%v{1,}]..$|^.[^%v].[%v{1,}].$|^.[^%v]..[%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-1].Character != char) && (gridData[i+1].Character != char) && (gridData[i+2].Character != char) && (gridData[i+3].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			greyRegex = fmt.Sprintf(".[^%v]...", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 2, 7, 12, 17, 22, 27:
		//Character in third position
		greenRegex = fmt.Sprintf("..%v..", char)
		yellowRegex = fmt.Sprintf("[%v{1,}].[^%v]..$|^.[%v{1,}][^%v]..$|^..[^%v][%v{1,}].$|^..[^%v].[%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-2].Character != char) && (gridData[i-1].Character != char) && (gridData[i+1].Character != char) && (gridData[i+2].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			greyRegex = fmt.Sprintf("..[^%v]..", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 3, 8, 13, 18, 23, 28:
		//Character in fourth position
		greenRegex = fmt.Sprintf("...%v.", char)
		yellowRegex = fmt.Sprintf("[%v{1,}]..[^%v].$|^.[%v{1,}].[^%v].$|^..[%v{1,}][^%v].$|^...[^%v][%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-3].Character != char) && (gridData[i-2].Character != char) && (gridData[i-1].Character != char) && (gridData[i+1].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			greyRegex = fmt.Sprintf("...[^%v].", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 4, 9, 14, 19, 24, 29:
		//Character in fifth position
		greenRegex = fmt.Sprintf("....%v", char)
		yellowRegex = fmt.Sprintf("[%v{1,}]...[^%v]$|^.[%v{1,}]..[^%v]$|^..[%v{1,}].[^%v]$|^...[%v{1,}][^%v]", char, char, char, char, char, char, char, char)
		if (gridData[i-4].Character != char) && (gridData[i-3].Character != char) && (gridData[i-2].Character != char) && (gridData[i-1].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			greyRegex = fmt.Sprintf("....[^%v]", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}
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

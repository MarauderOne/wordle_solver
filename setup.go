package main

import (
	"fmt"
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"github.com/golang/glog"
	"strings"
)

// Define a data structure for user inputs
type BoxData struct {
	Character string `json:"character"`
	Color     string `json:"color"`
}

// Define a function to initialise the answerList
func createNewAnswerList() (dictionary *dictionary_tools.MySimpleDict) {
	glog.Info("Creating d variable as NewSimpleDict struct")
	d := dictionary_tools.NewSimpleDict()
	glog.Info("Loading dictionary_tools/initialList.dict into d variable")
	d.Load("dictionary_tools/initialList.dict")
	glog.Info("Returning createNewAnswerList function")
	return d
}

// Define a function to set regex lookup patterns
func setRegexPatterns(i int, char string, gridData []BoxData) (greenRegex, yellowRegex, greyRegex string) {
	//Assign position based on index number in the array (i)
	switch i {
	case 0, 5, 10, 15, 20, 25:
		//Character in first position
		glog.Info("Set greenRegex pattern for the first position")
		greenRegex = fmt.Sprintf("%v....", char)
		glog.Info("Set yellowRegex pattern for the first position")
		yellowRegex = fmt.Sprintf("[^%v][%v{1,}]...$|^[^%v].[%v{1,}]..$|^[^%v]..[%v{1,}].$|^[^%v]...[%v{1,}]", char, char, char, char, char, char, char, char)

		if (gridData[i+1].Character != char) && (gridData[i+2].Character != char) && (gridData[i+3].Character != char) && (gridData[i+4].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			glog.Info("Set greyRegex pattern for the first position (where there are no matching characters in the rest of the word)")
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			glog.Info("Set greyRegex pattern for the first position (exclude words with this character in this position only)")
			greyRegex = fmt.Sprintf("[^%v]....", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 1, 6, 11, 16, 21, 26:
		//Character in second position
		glog.Info("Set greenRegex pattern for the second position")
		greenRegex = fmt.Sprintf(".%v...", char)
		glog.Info("Set yellowRegex pattern for the second position")
		yellowRegex = fmt.Sprintf("[%v{1,}][^%v]...$|^.[^%v][%v{1,}]..$|^.[^%v].[%v{1,}].$|^.[^%v]..[%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-1].Character != char) && (gridData[i+1].Character != char) && (gridData[i+2].Character != char) && (gridData[i+3].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			glog.Info("Set greyRegex pattern for the second position (where there are no matching characters in the rest of the word)")
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			glog.Info("Set greyRegex pattern for the second position (exclude words with this character in this position only)")
			greyRegex = fmt.Sprintf(".[^%v]...", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 2, 7, 12, 17, 22, 27:
		//Character in third position
		glog.Info("Set greenRegex pattern for the third position")
		greenRegex = fmt.Sprintf("..%v..", char)
		glog.Info("Set yellowRegex pattern for the third position")
		yellowRegex = fmt.Sprintf("[%v{1,}].[^%v]..$|^.[%v{1,}][^%v]..$|^..[^%v][%v{1,}].$|^..[^%v].[%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-2].Character != char) && (gridData[i-1].Character != char) && (gridData[i+1].Character != char) && (gridData[i+2].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			glog.Info("Set greyRegex pattern for the third position (where there are no matching characters in the rest of the word)")
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			glog.Info("Set greyRegex pattern for the third position (exclude words with this character in this position only)")
			greyRegex = fmt.Sprintf("..[^%v]..", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 3, 8, 13, 18, 23, 28:
		//Character in fourth position
		glog.Info("Set greenRegex pattern for the fourth position")
		greenRegex = fmt.Sprintf("...%v.", char)
		glog.Info("Set yellowRegex pattern for the fourth position")
		yellowRegex = fmt.Sprintf("[%v{1,}]..[^%v].$|^.[%v{1,}].[^%v].$|^..[%v{1,}][^%v].$|^...[^%v][%v{1,}]", char, char, char, char, char, char, char, char)
		if (gridData[i-3].Character != char) && (gridData[i-2].Character != char) && (gridData[i-1].Character != char) && (gridData[i+1].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			glog.Info("Set greyRegex pattern for the fourth position (where there are no matching characters in the rest of the word)")
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			glog.Info("Set greyRegex pattern for the fourth position (exclude words with this character in this position only)")
			greyRegex = fmt.Sprintf("...[^%v].", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}

	case 4, 9, 14, 19, 24, 29:
		//Character in fifth position
		glog.Info("Set greenRegex pattern for the fifth position")
		greenRegex = fmt.Sprintf("....%v", char)
		glog.Info("Set yellowRegex pattern for the fifth position")
		yellowRegex = fmt.Sprintf("[%v{1,}]...[^%v]$|^.[%v{1,}]..[^%v]$|^..[%v{1,}].[^%v]$|^...[%v{1,}][^%v]", char, char, char, char, char, char, char, char)
		if (gridData[i-4].Character != char) && (gridData[i-3].Character != char) && (gridData[i-2].Character != char) && (gridData[i-1].Character != char) {
			//If the letter is grey and does not exist in any other position in the word then we can eliminate words which contain that letter anywhere
			glog.Info("Set greyRegex pattern for the fifth position (where there are no matching characters in the rest of the word)")
			greyRegex = fmt.Sprintf("[^%v]*", char)
		} else {
			//Just eliminate words with this character in the current position
			glog.Info("Set greyRegex pattern for the fifth position (exclude words with this character in this position only)")
			greyRegex = fmt.Sprintf("....[^%v]", char)
			//This solution is not perfect as it does not account for cases where there is a grey "A" (for example), and one or more green "A"s
			//ToDo: Implement an efficient method of determining which positions have common characters, then write a regex patter to keep words for the green positions and delete any other occurrences of that character
		}
	}
	glog.Info("Returning setRegexPatterns function")
	return greenRegex, yellowRegex, greyRegex
}

// Define a function to revise the answerList based on given regex patterns
func reviseAnswerList(answersList *dictionary_tools.MySimpleDict, regexPattern string) (revisedDictionary *dictionary_tools.MySimpleDict) {
	glog.Info("Filtering list of potential answers using regex to create new list of potential answers in newAnswerList variable")
	newAnswerList := answersList.Lookup(regexPattern, 0, 6000)
	glog.Info("Creating d variable as NewSimpleDict struct")
	d := dictionary_tools.NewSimpleDict()
	glog.Info("Loading newAnswerList variable into d variable")
	d.AddWordsList(newAnswerList)
	return d
}

// Define a function to determine if a character is not alphabetic
func nonAlpha(char string) bool {

	//Define a list of alphabetic characters
	glog.Info("Defining list of alphabetic characters")
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	glog.Info("Checking if box character is alphabetic")
	if (strings.Contains(alpha, strings.ToUpper(string(char)))) || (string(char) == "") {
		glog.Info("Box character is alphabetic")
		return false
	} else {
		glog.Errorf("Box character is not alphabetic: %v", char)
		return true
	}
}

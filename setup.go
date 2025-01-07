package main

import (
	"fmt"
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"github.com/golang/glog"
	"strings"
)

//Define a data structure for user inputs
type CellData struct {
	Character string `json:"character"`
	Color     string `json:"color"`
}

//Define a function to initialise the answerList
func createNewAnswerList() (dictionary *dictionary_tools.MySimpleDict) {
	glog.Info("Creating d variable as NewSimpleDict struct")
	newDictionary := dictionary_tools.NewSimpleDict()
	glog.Info("Loading dictionary_tools/initialList.dict into d variable")
	newDictionary.Load("dictionary_tools/initialList.dict")
	glog.Info("Returning createNewAnswerList function")
	return newDictionary
}

//Define a function to set a regex pattern for the Lookup method used in the reviseAnswerList function
//Specifically, these regex patterns relate to the character & color of the supplied box without regard for the rest of the word
func setSingleLetterRegexPattern(i int, char string, gridData []CellData) (singleLetterRegexPattern string) {
	//Assign position based on index number in the array (i)
	switch i {
	case 0, 5, 10, 15, 20, 25:
		//Character in first position
		if (gridData[i].Color == "green") {
			//Character is green
			glog.Info("Set green singleLetterRegexPattern for the first position")
			singleLetterRegexPattern = fmt.Sprintf("%v....", char)
		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			glog.Info("Set yellow singleLetterRegexPattern for the first position")
			singleLetterRegexPattern = fmt.Sprintf("[^%v].*[%v{1,4}].*", char, char)
		} else {
			//Character is grey
			glog.Info("Set grey singleLetterRegexPattern for the first position")
			singleLetterRegexPattern = fmt.Sprintf("[^%v]....", char)
		}

	case 1, 6, 11, 16, 21, 26:
		//Character in second position
		if (gridData[i].Color == "green") {
			//Character is green
			glog.Info("Set green singleLetterRegexPattern for the second position")
			singleLetterRegexPattern = fmt.Sprintf(".%v...", char)
		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			glog.Info("Set yellow singleLetterRegexPattern for the second position")
			singleLetterRegexPattern = fmt.Sprintf("%v[^%v]...$|^.[^%v].*[%v{1,3}].*", char, char, char, char)
		} else {
			//Character is grey
			glog.Info("Set grey singleLetterRegexPattern for the second position")
			singleLetterRegexPattern = fmt.Sprintf(".[^%v]...", char)
		}

	case 2, 7, 12, 17, 22, 27:
		//Character in third position
		if (gridData[i].Color == "green") {
			//Character is green
			glog.Info("Set green singleLetterRegexPattern for the third position")
			singleLetterRegexPattern = fmt.Sprintf("..%v..", char)
		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			glog.Info("Set yellow singleLetterRegexPattern for the third position")
			singleLetterRegexPattern = fmt.Sprintf(".*[%v{1,2}].*[^%v]..$|^..[^%v].*[%v{1,2}].*", char, char, char, char)
		} else {
			//Character is grey
			glog.Info("Set grey singleLetterRegexPattern for the third position")
			singleLetterRegexPattern = fmt.Sprintf("..[^%v]..", char)
		}

	case 3, 8, 13, 18, 23, 28:
		//Character in fourth position
		if (gridData[i].Color == "green") {
			//Character is green
			glog.Info("Set green singleLetterRegexPattern for the fourth position")
			singleLetterRegexPattern = fmt.Sprintf("...%v.", char)
		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			glog.Info("Set yellow singleLetterRegexPattern for the fourth position")
			singleLetterRegexPattern = fmt.Sprintf(".*[%v{1,3}].*[^%v].$|^...[^%v]%v", char, char, char, char)
		} else {
			//Character is grey
			glog.Info("Set grey singleLetterRegexPattern for the fourth position")
			singleLetterRegexPattern = fmt.Sprintf("...[^%v].", char)
		}

	case 4, 9, 14, 19, 24, 29:
		//Character in fifth position
		if (gridData[i].Color == "green") {
			//Character is green
			glog.Info("Set green singleLetterRegexPattern for the fourth position")
			singleLetterRegexPattern = fmt.Sprintf("....%v", char)
		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			glog.Info("Set yellow singleLetterRegexPattern for the fifth position")
			singleLetterRegexPattern = fmt.Sprintf(".*[%v{1,4}].*[^%v]", char, char)
		} else {
			//Character is grey
			glog.Info("Set grey singleLetterRegexPattern for the fifth position")
			singleLetterRegexPattern = fmt.Sprintf("....[^%v]", char)
		}
	}
	glog.Info("Returning setRegexPattern function")
	return singleLetterRegexPattern
}

//Define a function to set regex patterns for the Lookup method used in the reviseAnswerList function
//Specifically, these regex patterns relate to how the relationships between different colors of matching characters can tells us things about the whole word (i.e. more complicated logic)
func setMultiLetterRegexPattern(i int, char string, gridData []CellData) (multiLetterRegexPattern string) {

	//Initialise array, so that we can return multiple patterns as a single variable
	regexPatternArray := []string{}

	//Assign position based on index number in the array (i)
	switch i {
	case 0, 5, 10, 15, 20, 25:
		//Character in first position
		if (gridData[i].Color == "green") {
			//Character is green

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is green with matching grey in the second position
				glog.Info("Set regex pattern for first position green character with matching grey character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]...", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is green with matching grey in the third position
				glog.Info("Set regex pattern for first position green character with matching grey character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v.[^%v]..", char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) {
				//Character is green with matching grey in the fourth position
				glog.Info("Set regex pattern for first position green character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v..[^%v].", char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "grey")) {
				//Character is green with matching grey in the fifth position
				glog.Info("Set regex pattern for first position green character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v...[^%v]", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is green with matching yellow in the second position
				glog.Info("Set regex pattern for first position green character with matching yellow character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v].*[%v{1,3}].*", char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is green with matching yellow in the third position
				glog.Info("Set regex pattern for first position green character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v[^%v]..$|^%v[^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is green with matching yellow in the fourth position
				glog.Info("Set regex pattern for first position green character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v.*[%v{1,2}].*[^%v].$|^%v..[^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "yellow")) {
				//Character is green with matching yellow in the fifth position
				glog.Info("Set regex pattern for first position green character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v.*[%v{1,4}].*[^%v]", char, char, char))
			}

		} else if (gridData[i].Color == "yellow") {
			//Character is yellow
			
			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is yellow with matching green in the second position
				glog.Info("Set regex pattern for first position yellow character with matching green character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v.*[%v{1,3}].*", char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is yellow with matching green in the third position
				glog.Info("Set regex pattern for first position yellow character with matching green character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v%v..$|^[^%v].%v.*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "green")) {
				//Character is yellow with matching green in the fourth position
				glog.Info("Set regex pattern for first position yellow character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,2}].*%v.$|^[^%v]..%v%v", char, char, char, char, char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "green")) {
				//Character is yellow with matching green in the fifth position
				glog.Info("Set regex pattern for first position yellow character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,3}].*%v", char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is yellow with matching yellow in the second position
				glog.Info("Set regex pattern for first position yellow character with matching yellow character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v][^%v].*[%v{2,3}].*", char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is yellow with matching yellow in the third position
				glog.Info("Set regex pattern for first position yellow character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v[^%v].*[%v{1,2}].*$|^[^%v].[^%v]%v%v", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is yellow with matching yellow in the fourth position
				glog.Info("Set regex pattern for first position yellow character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v%v[^%v].$|^[^%v].*[%v{1,2}].*[^%v]%v", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "yellow")) {
				//Character is yellow with matching yellow in the fifth position
				glog.Info("Set regex pattern for first position yellow character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{2,3}].*[^%v]", char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is yellow with matching grey in the second position
				glog.Info("Set regex pattern for first position yellow character with matching grey character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v][^%v].*[%v{1,3}].*", char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is yellow with matching grey in the third position
				glog.Info("Set regex pattern for first position yellow character with matching grey character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v[^%v]..$|^[^%v].[^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) {
				//Character is yellow with matching grey in the fourth position
				glog.Info("Set regex pattern for first position yellow character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,2}].*[^%v].$|^[^%v]..[^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "grey")) {
				//Character is yellow with matching grey in the fifth position
				glog.Info("Set regex pattern for first position yellow character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,3}].*[^%v]", char, char, char))
			}

		} else {
			//Character is grey

			if ((gridData[i+1].Character != char) || (gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) && ((gridData[i+2].Character != char) || (gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) && ((gridData[i+3].Character != char) || (gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) && ((gridData[i+4].Character != char) || (gridData[i+4].Character == char) && (gridData[i+4].Color == "grey")) {
				//Character is grey with no matching green or yellow characters in the rest of the word
				glog.Info("Set regex pattern for first position grey character with no matching green or yellow characters in the rest of the word")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]*", char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is grey with matching green in the second position
				glog.Info("Set regex pattern for first position grey character with matching green character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v...", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is grey with matching green in the third position
				glog.Info("Set regex pattern for first position grey character with matching green character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].%v..", char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "green")) {
				//Character is grey with matching green in the fourth position
				glog.Info("Set regex pattern for first position grey character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]..%v.", char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "green")) {
				//Character is grey with matching green in the fifth position
				glog.Info("Set regex pattern for first position grey character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]...%v", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is grey with matching yellow in the second position
				glog.Info("Set regex pattern for first position grey character with matching yellow character in the second position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v][^%v].*[%v{1,3}].*", char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is grey with matching yellow in the third position
				glog.Info("Set regex pattern for first position grey character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]%v[^%v]..$|^[^%v].[^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is grey with matching yellow in the fourth position
				glog.Info("Set regex pattern for first position grey character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,2}].*[^%v].$|^[^%v]..[^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+4].Character == char) && (gridData[i+4].Color == "yellow")) {
				//Character is grey with matching yellow in the fifth position
				glog.Info("Set regex pattern for first position grey character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v].*[%v{1,3}].*[^%v]", char, char, char))
			}
		}

	case 1, 6, 11, 16, 21, 26:
		//Character in second position
		if (gridData[i].Color == "green") {
			//Character is green

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is green with matching grey in the third position
				glog.Info("Set regex pattern for second position green character with matching grey character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".%v[^%v]..", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is green with matching grey in the fourth position
				glog.Info("Set regex pattern for second position green character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".%v.[^%v].", char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) {
				//Character is green with matching grey in the fifth position
				glog.Info("Set regex pattern for second position green character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".%v..[^%v]", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is green with matching yellow in the third position
				glog.Info("Set regex pattern for second position green character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v[^%v]..$|^.%v[^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is green with matching yellow in the fourth position
				glog.Info("Set regex pattern for second position green character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v.[^%v].$|^.%v%v.[^%v].$|^.%v.[^%v]%v", char, char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is green with matching yellow in the fifth position
				glog.Info("Set regex pattern for second position green character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v..[^%v]$|^.%v.*[%v{1,2}].*[^%v]", char, char, char, char, char, char))
			}

		} else if (gridData[i].Color == "yellow") {
			//Character is yellow

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is yellow with matching green in the third position
				glog.Info("Set regex pattern for second position yellow character with matching green character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]%v..$|^.[^%v]%v.*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is yellow with matching green in the fourth position
				glog.Info("Set regex pattern for second position yellow character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v].%v.$|^.[^%v]%v%v.$|^.[^%v].%v%v", char, char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "green")) {
				//Character is yellow with matching green in the fifth position
				glog.Info("Set regex pattern for second position yellow character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]..%v$|^.[^%v].*[%v{1,2}].*%v", char, char, char, char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is yellow with matching yellow in the third position
				glog.Info("Set regex pattern for second position yellow character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v][^%v].*[%v{1,2}].*$|^.[^%v][^%v]%v%v", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is yellow with matching yellow in the fourth position
				glog.Info("Set regex pattern for second position yellow character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]%v[^%v].$|^%v[^%v].[^%v]%v$|^.[^%v]%v[^%v]%v", char, char, char, char, char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is yellow with matching yellow in the fifth position
				glog.Info("Set regex pattern for second position yellow character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v].*[%v{1,2}].*[^%v]$|^.[^%v]%v%v[^%v]", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is yellow with matching grey in the third position
				glog.Info("Set regex pattern for second position yellow character with matching grey character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v][^%v]..$|^.[^%v][^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is yellow with matching grey in the fourth position
				glog.Info("Set regex pattern for second position yellow character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v].[^%v].$|^.[^%v]%v[^%v].$|^.[^%v].[^%v]%v", char, char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) {
				//Character is yellow with matching grey in the fifth position
				glog.Info("Set regex pattern for second position yellow character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]..[^%v]$|^.[^%v].*[%v{1,2}].*[^%v]", char, char, char, char, char, char))
			}

		} else {
			//Character is grey

			if ((gridData[i-1].Character != char) || (gridData[i-1].Character == char) && (gridData[i-1].Color == "grey")) && ((gridData[i+1].Character != char) || (gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) && ((gridData[i+2].Character != char) || (gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) && ((gridData[i+3].Character != char) || (gridData[i+3].Character == char) && (gridData[i+3].Color == "grey")) {
				//Character is grey with no matching green or yellow characters in the rest of the word
				glog.Info("Set regex pattern for second position grey character with no matching green or yellow characters in the rest of the word")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]*", char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is grey with matching green in the third position
				glog.Info("Set regex pattern for second position grey character with matching green character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".[^%v]%v..", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is grey with matching green in the fourth position
				glog.Info("Set regex pattern for second position grey character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".[^%v].%v.", char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "green")) {
				//Character is grey with matching green in the fifth position
				glog.Info("Set regex pattern for second position grey character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".[^%v]..%v", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is grey with matching yellow in the third position
				glog.Info("Set regex pattern for second position grey character with matching yellow character in the third position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v][^%v]..$|^.[^%v][^%v].*[%v{1,2}].*", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is grey with matching yellow in the fourth position
				glog.Info("Set regex pattern for second position grey character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v].[^%v].$|^.[^%v]%v[^%v].$|^.[^%v].[^%v]%v", char, char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+3].Character == char) && (gridData[i+3].Color == "yellow")) {
				//Character is grey with matching yellow in the fifth position
				glog.Info("Set regex pattern for second position grey character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v[^%v]..[^%v]$|^.[^%v].*[%v{1,2}].*[^%v]", char, char, char, char, char, char))
			}
		}

	case 2, 7, 12, 17, 22, 27:
		//Character in third position
		if (gridData[i].Color == "green") {
			//Character is green

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is green with matching grey in the fourth position
				glog.Info("Set regex pattern for third position green character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("..%v[^%v].", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is green with matching grey in the fifth position
				glog.Info("Set regex pattern for third position green character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("..%v.[^%v]", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is green with matching yellow in the fourth position
				glog.Info("Set regex pattern for third position green character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*%v[^%v].$|^..%v[^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is green with matching yellow in the fifth position
				glog.Info("Set regex pattern for third position green character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*%v.[^%v]$|^..%v%v[^%v]", char, char, char, char, char, char))
			}

		} else if (gridData[i].Color == "yellow") {
			//Character is yellow

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is yellow with matching green in the fourth position
				glog.Info("Set regex pattern for third position yellow character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v]%v.$|^..[^%v]%v%v", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is yellow with matching green in the fifth position
				glog.Info("Set regex pattern for third position yellow character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v].%v$|^..[^%v]%v%v", char, char, char, char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is yellow with matching yellow in the fourth position
				glog.Info("Set regex pattern for third position yellow character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v[^%v][^%v].$|^.*[%v{1,2}].*[^%v][^%v]%v", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is yellow with matching yellow in the fifth position
				glog.Info("Set regex pattern for third position yellow character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("%v%v[^%v].[^%v]$|^.*[%v{1,2}].*[^%v]%v[^%v]", char, char, char, char, char, char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is yellow with matching grey in the fourth position
				glog.Info("Set regex pattern for third position yellow character with matching grey character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v][^%v].$|^..[^%v][^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is yellow with matching grey in the fifth position
				glog.Info("Set regex pattern for third position yellow character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v].[^%v]$|^..[^%v]%v[^%v]", char, char, char, char, char, char))
			}

		} else {
			//Character is grey

			if ((gridData[i-2].Character != char) || (gridData[i-2].Character == char) && (gridData[i-2].Color == "grey")) && ((gridData[i-1].Character != char) || (gridData[i-1].Character == char) && (gridData[i-1].Color == "grey")) && ((gridData[i+1].Character != char) || (gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) && ((gridData[i+2].Character != char) || (gridData[i+2].Character == char) && (gridData[i+2].Color == "grey")) {
				//Character is grey with no matching green or yellow characters in the rest of the word
				glog.Info("Set regex pattern for third position grey character with no matching green or yellow characters in the rest of the word")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]*", char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is grey with matching green in the fourth position
				glog.Info("Set regex pattern for third position grey character with matching green character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("..[^%v]%v.", char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "green")) {
				//Character is grey with matching green in the fifth position
				glog.Info("Set regex pattern for third position grey character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("..[^%v].%v", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is grey with matching yellow in the fourth position
				glog.Info("Set regex pattern for third position grey character with matching yellow character in the fourth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v][^%v].$|^..[^%v][^%v]%v", char, char, char, char, char, char))
			}

			if ((gridData[i+2].Character == char) && (gridData[i+2].Color == "yellow")) {
				//Character is grey with matching yellow in the fifth position
				glog.Info("Set regex pattern for third position grey character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,2}].*[^%v].[^%v]$|^..[^%v]%v[^%v]", char, char, char, char, char, char))
			}

		}

	case 3, 8, 13, 18, 23, 28:
		//Character in fourth position
		if (gridData[i].Color == "green") {
			//Character is green

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is green with matching grey in the fifth position
				glog.Info("Set regex pattern for fourth position green character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("...%v[^%v]", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is green with matching yellow in the fifth position
				glog.Info("Set regex pattern for fourth position green character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,3}].*%v[^%v]", char, char, char))
			}

		} else if (gridData[i].Color == "yellow") {
			//Character is yellow

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is yellow with matching green in the fifth position
				glog.Info("Set regex pattern for fourth position yellow character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,3}].*[^%v]%v", char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is yellow with matching yellow in the fifth position
				glog.Info("Set regex pattern for fourth position yellow character with matching yellow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{2,3}].*[^%v][^%v]", char, char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is yellow with matching grey in the fifth position
				glog.Info("Set regex pattern for fourth position yellow character with matching grey character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,3}].*[^%v][^%v]", char, char, char))
			}

		} else {
			//Character is grey

			if ((gridData[i-3].Character != char) || (gridData[i-3].Character == char) && (gridData[i-3].Color == "grey")) && ((gridData[i-2].Character != char) || (gridData[i-2].Character == char) && (gridData[i-2].Color == "grey")) && ((gridData[i-1].Character != char) || (gridData[i-1].Character == char) && (gridData[i-1].Color == "grey")) && ((gridData[i+1].Character != char) || (gridData[i+1].Character == char) && (gridData[i+1].Color == "grey")) {
				//Character is grey with no matching green or yellow characters in the rest of the word
				glog.Info("Set regex pattern for fourth position grey character with no matching green or yellow characters in the rest of the word")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]*", char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "green")) {
				//Character is grey with matching green in the fifth position
				glog.Info("Set regex pattern for fourth position grey character with matching green character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("...[^%v]%v", char, char))
			}

			if ((gridData[i+1].Character == char) && (gridData[i+1].Color == "yellow")) {
				//Character is grey with matching yellow in the fifth position
				glog.Info("Set regex pattern for fourth position grey character with matching yelllow character in the fifth position")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf(".*[%v{1,3}].*[^%v][^%v]", char, char, char))
			}

		}

	case 4, 9, 14, 19, 24, 29:
		//Character in fifth position
		if (gridData[i].Color == "green") {
			//Character is green


		} else if (gridData[i].Color == "yellow") {
			//Character is yellow


		} else {
			//Character is grey

			if ((gridData[i-4].Character != char) || (gridData[i-4].Character == char) && (gridData[i-4].Color == "grey")) && ((gridData[i-3].Character != char) || (gridData[i-3].Character == char) && (gridData[i-3].Color == "grey")) && ((gridData[i-2].Character != char) || (gridData[i-2].Character == char) && (gridData[i-2].Color == "grey")) && ((gridData[i-1].Character != char) || (gridData[i-1].Character == char) && (gridData[i-1].Color == "grey")) {
				//Character is grey with no matching green or yellow characters in the rest of the word
				glog.Info("Set regex pattern for fifth position grey character with no matching green or yellow characters in the rest of the word")
				regexPatternArray = append(regexPatternArray, fmt.Sprintf("[^%v]*", char))
			}
		}
	}

	//Joining various regex patterns with separator
	glog.Info("Joining various regex patterns")
	multiLetterRegexPattern = strings.Join(regexPatternArray, "$|^")

	glog.Info("Returning setRegexPattern function")
	return multiLetterRegexPattern
}

//Define a function to revise the answerList based on given regex patterns
func reviseAnswerList(answersList *dictionary_tools.MySimpleDict, regexPattern string) (revisedDictionary *dictionary_tools.MySimpleDict) {
	glog.Info("Filtering list of potential answers using regex to create new list of potential answers in newAnswerList variable")
	newAnswerList := answersList.Lookup(regexPattern, 0, 6000)
	glog.Info("Creating d variable as NewSimpleDict struct")
	d := dictionary_tools.NewSimpleDict()
	glog.Info("Loading newAnswerList variable into d variable")
	d.AddWordsList(newAnswerList)
	return d
}

//Define a function to determine if a character is not an uppercase alphabetic character
func nonAlpha(char string) bool {

	//Define a list of uppercase alphabetic characters
	glog.Info("Defining list of uppercase alphabetic characters")
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//If the character in the box is in the alpha string or is ""
	glog.Info("Checking if box character is alphabetic")
	if (strings.Contains(alpha, strings.ToUpper(string(char)))) || (string(char) == "") {
		//Then return no error
		glog.Info("Box character is alphabetic")
		return false
	} else {
		//Else return an error
		glog.Errorf("Box character is not alphabetic: %v", char)
		return true
	}
}

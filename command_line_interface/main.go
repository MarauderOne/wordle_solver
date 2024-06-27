package main

import (
	"bufio"
	"fmt"
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Initialise answer list
	answerList := createNewAnswerList()

	//Ask user which letters that have ruled out
	negativeAnswers := askUserAboutNegatives()
	//Turn that answer into a regex pattern
	regexPattern := createNegativeRegexPattern(negativeAnswers)
	//Revise the answerList using the regex pattern
	answerList = reviseAnswerList(answerList, regexPattern)
	//Tell the user what the number of possible answers has been reduced to
	fmt.Printf("\nThe list of possible answers has been reduced to: %v\n", answerList.Count())
	fmt.Println()

	for _, position := range positions {
		//Ask user which letter they have found for the current position in the loop
		positiveAnswers := askUserAboutPositives(position)
		//Turn that answer into a regex pattern
		regexPattern := positiveAnswers.createRegexPattern()
		//Revise the answerList using the regex pattern
		answerList = reviseAnswerList(answerList, regexPattern)
		//Tell the user what the number of possible answers has been reduced to
		fmt.Printf("\nThe list of possible answers has been reduced to: %v\n", answerList.Count())
		//Tell the user what the list of possible answers has been reduced to
		output := strings.Join(answerList.Words, " ")
		fmt.Println(output)
		fmt.Println()
		//If we're down to a single possible answer, break the loop
		if answerList.Count() == 1 {
			break
		}
	}

}

// Define a list of possible letter positions
var positions = []string{"first", "second", "third", "fourth", "fifth"}

// Define a data struct for each positive answer
type posAnswer struct {
	character     string
	position      string
	knownPosition bool
}

// Define a function to initialise the answerList
func createNewAnswerList() (dictionary *dictionary_tools.MySimpleDict) {
	d := dictionary_tools.NewSimpleDict()
	d.Load("../dictionary_tools/initialList.dict")
	return d
}

// Define a function to ask the user which characters they have been able to rule out
func askUserAboutNegatives() (answer string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Which letters have you been able to rule out? Please type all characters on a single line.\n")
	answer, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//ToDo: Some sort of character validation
	return answer
}

// Define a function to ask the user which characters they have found and whether they are green/yellow
func askUserAboutPositives(position string) (answers posAnswer) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the letter in the %v position? (Enter * for unknown.)\n", position)
	characterAnswer, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	characterAnswerCleaned := strings.Trim(characterAnswer, "\n")
	//ToDo: Some sort of character validation

	if characterAnswerCleaned == "*" {
		return posAnswer{
			character:     characterAnswerCleaned,
			position:      position,
			knownPosition: false,
		}
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Is the position of the character confirmed? (Green = true OR Yellow = false)\n")
	knownPositionAnswer, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	knownPositionAnswerCleaned := strings.Trim(knownPositionAnswer, "\n")
	//ToDo: Some sort of character validation
	knownPositionAsBoolean, err := strconv.ParseBool(knownPositionAnswerCleaned)
	if err != nil {
		panic(err)
	}

	return posAnswer{
		character:     characterAnswerCleaned,
		position:      position,
		knownPosition: knownPositionAsBoolean,
	}
}

// Define a function to turn the negativeAnswers into a regex pattern
func createNegativeRegexPattern(negativeAnswers string) (regex string) {
	r := fmt.Sprintf("[^%v][^%v][^%v][^%v][^%v]", negativeAnswers, negativeAnswers, negativeAnswers, negativeAnswers, negativeAnswers)
	return r
}

// Define a method that is applied to answerStruct in order to determine the appropriate regex pattern
func (answerStruct *posAnswer) createRegexPattern() (regexPattern string) {
	if answerStruct.character == "*" {
		return "....."
	}

	switch answerStruct.position {
	case "first":
		if answerStruct.knownPosition {
			kpc := fmt.Sprintf("%v....", answerStruct.character)
			return kpc
		} else {
			upc := fmt.Sprintf("[^%v][%v{1,}]...$|^[^%v].[%v{1,}]..$|^[^%v]..[%v{1,}].$|^[^%v]...[%v{1,}]", answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character)
			return upc
		}
	case "second":
		if answerStruct.knownPosition {
			kpc := fmt.Sprintf(".%v...", answerStruct.character)
			return kpc
		} else {
			upc := fmt.Sprintf("[%v{1,}][^%v]...$|^.[^%v][%v{1,}]..$|^.[^%v].[%v{1,}].$|^.[^%v]..[%v{1,}]", answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character)
			return upc
		}
	case "third":
		if answerStruct.knownPosition {
			kpc := fmt.Sprintf("..%v..", answerStruct.character)
			return kpc
		} else {
			upc := fmt.Sprintf("[%v{1,}].[^%v]..$|^.[%v{1,}][^%v]..$|^..[^%v][%v{1,}].$|^..[^%v].[%v{1,}]", answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character)
			return upc
		}
	case "fourth":
		if answerStruct.knownPosition {
			kpc := fmt.Sprintf("...%v.", answerStruct.character)
			return kpc
		} else {
			upc := fmt.Sprintf("[%v{1,}]..[^%v].$|^.[%v{1,}].[^%v].$|^..[%v{1,}][^%v].$|^...[^%v][%v{1,}]", answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character)
			return upc
		}
	case "fifth":
		if answerStruct.knownPosition {
			kpc := fmt.Sprintf("....%v", answerStruct.character)
			return kpc
		} else {
			upc := fmt.Sprintf("[%v{1,}]...[^%v]$|^.[%v{1,}]..[^%v]$|^...[%v{1,}].[^%v]$|^...[%v{1,}][^%v]", answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character, answerStruct.character)
			return upc
		}
	}
	return ""
}

// Define a function to revise the answerList based on given regex patterns
func reviseAnswerList(answersList *dictionary_tools.MySimpleDict, regexPattern string) (revisedDictionary *dictionary_tools.MySimpleDict) {
	newAnswerList := answersList.Lookup(regexPattern, 0, 6000)
	d := dictionary_tools.NewSimpleDict()
	d.AddWordsList(newAnswerList)
	return d
}

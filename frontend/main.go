package main

import (
	"fmt"
	"net/http"
    "github.com/MarauderOne/wordle_solver/dictionary_tools"
	"github.com/gin-gonic/gin"
    "strings"
)

func main() {

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Define your routes
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Endpoint to handle Wordle solving
	r.POST("/solve", solveWordle)

	// Run the server
	r.Run(":8080")
}

type BoxData struct {
    Character string `json:"character"`
    Color     string `json:"color"`
}

func solveWordle(c *gin.Context) {
	var gridData []BoxData
	if err := c.ShouldBindJSON(&gridData); err != nil {
		    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		    return
	}

	// Call your Wordle solving function (implement this)
	result, countOfResults, solveError := solve(gridData)

    if solveError != "" {
        //Respond with a 500 error
        c.JSON(http.StatusInternalServerError, gin.H{"error": solveError})
    } else {
        // Respond with the result
	    c.JSON(http.StatusOK, gin.H{"result": result, "resultSummary": countOfResults})
    }
}



// Placeholder for your Wordle solving logic
func solve(guess []BoxData) (result, countOfResults string, solveError string) {

    //Initialise answer list
    answerList := createNewAnswerList()

    //Start looping through the user boxes
    revisionLoop: for i, box := range guess {

        //Check for non-alphabetic characters
        if nonAlpha(box.Character) {
            solveError = fmt.Sprintf("Invalid character: %v", box.Character)
            break revisionLoop
        }

        if (box.Character == "") || (box.Color == "") {
            //Skip over boxes which have either no character or no color
            continue
            //ToDo: Add logging
        }

        var greenRegex string
        var yellowRegex string
        var greyRegex string

        //Assign postition based on index number in the array (i)
        switch i {
        case 0, 5, 10, 15, 20, 25:
            //Character in first position
            greenRegex = fmt.Sprintf("%v....", box.Character)
            yellowRegex = fmt.Sprintf("[^%v][%v{1,}]...$|^[^%v].[%v{1,}]..$|^[^%v]..[%v{1,}].$|^[^%v]...[%v{1,}]", box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character)
            greyRegex = fmt.Sprintf("[^%v]....", box.Character)
        case 1, 6, 11, 16, 21, 26:
            //Character in second position
            greenRegex = fmt.Sprintf(".%v...", box.Character)
            yellowRegex = fmt.Sprintf("[%v{1,}][^%v]...$|^.[^%v][%v{1,}]..$|^.[^%v].[%v{1,}].$|^.[^%v]..[%v{1,}]", box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character)
            greyRegex = fmt.Sprintf(".[^%v]...", box.Character)
        case 2, 7, 12, 17, 22, 27:
            //Character in third position
            greenRegex = fmt.Sprintf("..%v..", box.Character)
            yellowRegex = fmt.Sprintf("[%v{1,}].[^%v]..$|^.[%v{1,}][^%v]..$|^..[^%v][%v{1,}].$|^..[^%v].[%v{1,}]", box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character)
            greyRegex = fmt.Sprintf("..[^%v]..", box.Character)
        case 3, 8, 13, 18, 23, 28:
            //Character in fourth position
            greenRegex = fmt.Sprintf("...%v.", box.Character)
            yellowRegex = fmt.Sprintf("[%v{1,}]..[^%v].$|^.[%v{1,}].[^%v].$|^..[%v{1,}][^%v].$|^...[^%v][%v{1,}]", box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character)
            greyRegex = fmt.Sprintf("...[^%v].", box.Character)
        case 4, 9, 14, 19, 24, 29:
            //Character in fifth position
            greenRegex = fmt.Sprintf("....%v", box.Character)
            yellowRegex = fmt.Sprintf("[%v{1,}]...[^%v]$|^.[%v{1,}]..[^%v]$|^..[%v{1,}].[^%v]$|^...[%v{1,}][^%v]", box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character, box.Character)
            greyRegex = fmt.Sprintf("....[^%v]", box.Character)
        }

        switch box.Color {
        case "green":
            //Find matches for character in position
		    answerList = reviseAnswerList(answerList, greenRegex)
        case "yellow":
            //Find matches for character not position
            answerList = reviseAnswerList(answerList, yellowRegex)
        case "grey":
            //Find matches for character not position
            answerList = reviseAnswerList(answerList, greyRegex)
        default:
            //Invalid color, this should never be reached
            solveError = fmt.Sprintf("Invalid color: %v", box.Color)
            break revisionLoop
        }

        //Break the loop if potential answers drop to 1 or fewer
        if answerList.Count() <= 1 {
            break
        }
    }

        var resultSummary string = fmt.Sprintf("Potential answers: %v\n", answerList.Count())
        results := strings.Join(answerList.Words, " ")
        return results, resultSummary, solveError
}

//Define a function to initialise the answerList
func createNewAnswerList() (dictionary *dictionary_tools.MySimpleDict) {
	d := dictionary_tools.NewSimpleDict()
	d.Load("../dictionary_tools/initialList.dict")
	return d
}

//Define a function to revise the answerList based on given regex patterns
func reviseAnswerList(answersList *dictionary_tools.MySimpleDict, regexPattern string) (revisedDictionary *dictionary_tools.MySimpleDict) {
	newAnswerList := answersList.Lookup(regexPattern, 0, 6000)
	d := dictionary_tools.NewSimpleDict()
	d.AddWordsList(newAnswerList)
	return d
}

const alpha = "abcdefghijklmnopqrstuvwxyz"
func nonAlpha(s string) bool {
    for _, char := range s {  
       if strings.Contains(alpha, strings.ToLower(string(char))) {
          return false
       }
    }
    return true
 }

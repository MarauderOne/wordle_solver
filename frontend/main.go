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
	result, countOfResults := solve(gridData)

	// Respond with the result
	c.JSON(http.StatusOK, gin.H{"result": result, "resultSummary": countOfResults})
}



// Placeholder for your Wordle solving logic
func solve(guess []BoxData) (result, countOfResults string) {

    //Initialise answer list
    answerList := createNewAnswerList()

    //Start looping through the user boxes
    for i, box := range guess {

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

        
        if box.Color == "green" {
            //Find matches for character in position
		    answerList = reviseAnswerList(answerList, greenRegex)

            } else if box.Color == "yellow" {
                //Find matches for character anywhere but in position
                answerList = reviseAnswerList(answerList, yellowRegex)

                } else if box.Color == "grey" {
                    //Find matches for character not position
                    answerList = reviseAnswerList(answerList, greyRegex)

                } else {
                //Invalid color, this should never be reached
        }

        //Break the loop if potential answers drop to 1 or fewer
        if answerList.Count() <= 1 {
            break
        }
    }

    //return fmt.Sprint("End of function")
    var resultSummary string = fmt.Sprintf("Potential answers: %v\n", answerList.Count())
    results := strings.Join(answerList.Words, " ")
    return results, resultSummary 
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

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
revisionLoop:
	for i, box := range guess {

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

		//Set regex patterns for each color according current index position
		greenRegex, yellowRegex, greyRegex := setRegexPatterns(i, box.Character)

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

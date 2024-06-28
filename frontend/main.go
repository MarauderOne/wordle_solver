package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
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

	//Define the port
	port := os.Getenv("PORT")
	//Define default port (for local testing)
	if port == "" {
		port = "8080"
	}

	// Run the server
	err := r.Run(":" + port)
	if err != nil {
		glog.Fatalf("Web server initialisation failed: %v", err)
	}
}

func solveWordle(c *gin.Context) {
	var gridData []BoxData
	if err := c.ShouldBindJSON(&gridData); err != nil {
		glog.Errorf("Unable to bind JSON from page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call your Wordle solving function (implement this)
	glog.Info("Calling solve function")
	result, countOfResults, solveError := solve(gridData)

	if solveError != "" {
		//Respond with a 500 error
		glog.Warningf("Responding to page with error response: %v", solveError)
		glog.Flush()
		c.JSON(http.StatusInternalServerError, gin.H{"error": solveError})
	} else {
		// Respond with the result
		glog.Info("Responding to the page with expected results")
		glog.Flush()
		c.JSON(http.StatusOK, gin.H{"result": result, "resultSummary": countOfResults})
	}
}

// Placeholder for your Wordle solving logic
func solve(gridData []BoxData) (result, countOfResults string, solveError string) {

	//Initialise answer list
	glog.Info("Calling createNewAnswerList function")
	answerList := createNewAnswerList()

	//Start looping through the user boxes
revisionLoop:
	for i, box := range gridData {

		//Check for non-alphabetic characters
		if nonAlpha(box.Character) {
			glog.Errorf("Invalid character recieved: %v", box.Character)
			solveError = fmt.Sprintf("Invalid character: %v", box.Character)
			glog.Error("Breaking revision loop")
			break revisionLoop
		}

		//Skip over boxes which have either no character or no color
		if (box.Character == "") || (box.Color == "") {
			glog.Info("Character or Color value is missing, skipping this iteration of revision loop")
			continue
		}

		//Set regex patterns for each color according current index position
		glog.Info("Calling setRegexPatterns function")
		greenRegex, yellowRegex, greyRegex := setRegexPatterns(i, box.Character, gridData)

		switch box.Color {
		case "green":
			//Find matches for character in position
			glog.Info("Calling reviseAnswerList function using greenRegex pattern")
			answerList = reviseAnswerList(answerList, greenRegex)
		case "yellow":
			//Find matches for character not position
			glog.Info("Calling reviseAnswerList function using yellowRegex pattern")
			answerList = reviseAnswerList(answerList, yellowRegex)
		case "grey":
			//Find matches for character not position
			glog.Info("Calling reviseAnswerList function using greyRegex pattern")
			answerList = reviseAnswerList(answerList, greyRegex)
		default:
			//Invalid color, this should never be reached
			glog.Errorf("Invalid color recieved: %v", box.Color)
			solveError = fmt.Sprintf("Invalid color: %v", box.Color)
			glog.Error("Breaking revision loop")
			break revisionLoop
		}

		//Break the loop if potential answers drop to 1 or fewer
		if answerList.Count() <= 1 {
			glog.Info("List of potential answers has reached 1 or fewer words, breaking revision loop")
			break
		}
	}

	glog.Info("Writing results summary")
	var resultSummary string = fmt.Sprintf("Potential answers: %v\n", answerList.Count())
	glog.Info("Writing results")
	results := strings.Join(answerList.Words, " ")
	glog.Info("Returning solve function")
	return results, resultSummary, solveError
}

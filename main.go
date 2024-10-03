package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {

	//Parse the argument flags (like the ones in Heroku's Procfile)
	flag.Parse()

	webServer := gin.Default()

	//Serve static files
	webServer.Static("/wordlesolver", "./frontend")

	//Define routes
	webServer.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	//Endpoint to handle Wordle solving
	webServer.POST("/guesses", parseGuesses)

	//Define the port
	port := os.Getenv("PORT")
	//Define default port (for local testing)
	if port == "" {
		port = "8080"
	}

	//Run the webserver
	err := webServer.Run(":" + port)
	if err != nil {
		glog.Fatalf("Web server initialisation failed: %v", err)
	}
}

//Define function to parse the user's guesses
func parseGuesses(c *gin.Context) {
	var gridData []CellData
	if err := c.ShouldBindJSON(&gridData); err != nil {
		glog.Errorf("Unable to bind JSON from page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Convert Character values to uppercase
	for i := range gridData {
			gridData[i].Character = strings.ToUpper(gridData[i].Character)
	}

	//Call the Wordle solving function
	glog.Infof("Calling solveWordle function with gridData: %v", gridData)
	result, countOfResults, solvingError, httpStatus := solveWordle(gridData)

	if solvingError != "" {
		//Respond with the http status code and error message returned by the solveWordle function
		glog.Warningf("Responding to page with error response: %v", solvingError)
		glog.Flush()
		c.JSON(httpStatus, gin.H{"error": solvingError})
	} else {
		//Respond with the result
		glog.Info("Responding to the page with expected results")
		glog.Flush()
		c.JSON(httpStatus, gin.H{"result": result, "resultCount": countOfResults})
	}
}

//Function for Wordle solving logic
func solveWordle(gridData []CellData) (result string, countOfResults int, solvingError string, httpStatus int) {

	//Set default HTTP response code (will be updated if there is an error)
	httpStatus = http.StatusOK

	//Initialise answer list
	glog.Info("Calling createNewAnswerList function")
	answerList := createNewAnswerList()

	//Start looping through the user boxes
	for i, box := range gridData {

		//Check for non-alphabetic characters
		if nonAlpha(box.Character) {
			glog.Errorf("Invalid character recieved: %v", box.Character)
			solvingError = fmt.Sprintf("Invalid character: %v", box.Character)
			httpStatus = http.StatusBadRequest
			glog.Error("Breaking revision loop")
			glog.Info("Writing results count")
			var resultCount int = answerList.Count()
			glog.Info("Writing results")
			results := strings.Join(answerList.Words, " ")
			glog.Info("Returning solveWordle function")
			return results, resultCount, solvingError, httpStatus
		}

		//Skip over boxes which have either no character or no color
		if (box.Character == "") || (box.Color == "") {
			glog.Info("Character or Color value is missing, skipping this iteration of revision loop")
			continue
		}

		//Set regex pattern for each box according current index position
		glog.Info("Calling setSingleLetterRegexPattern function")
		singleLetterRegexPattern := setSingleLetterRegexPattern(i, box.Character, gridData)
		glog.Info("Calling setMultiLetterRegexPattern function")
		multiLetterRegexPattern := setMultiLetterRegexPattern(i, box.Character, gridData)

		//Revise the answerList based on the box color
		switch box.Color {
		case "green", "yellow", "grey":
			//Find matches for character in position
			glog.Info("Calling reviseAnswerList function using singleLetterRegexPattern pattern")
			answerList = reviseAnswerList(answerList, singleLetterRegexPattern)

			//If a complex logic operator has been set, apply it to the list here
			if len(multiLetterRegexPattern) > 0 {
				glog.Info("Calling reviseAnswerList function using multiLetterRegexPattern pattern")
				answerList = reviseAnswerList(answerList, multiLetterRegexPattern)
			}
		default:
			//Invalid color, this should never be reached (except in the tests)
			glog.Errorf("Invalid color recieved: %v", box.Color)
			solvingError = fmt.Sprintf("Invalid color: %v", box.Color)
			httpStatus = http.StatusBadRequest
			glog.Error("Breaking revision loop")
			glog.Info("Writing results count")
			var resultCount int = answerList.Count()
			glog.Info("Writing results")
			results := strings.Join(answerList.Words, " ")
			glog.Info("Returning solveWordle function")
			return results, resultCount, solvingError, httpStatus
		}

		//Break the loop if potential answers drop to 1 or fewer
		if answerList.Count() <= 1 {
			glog.Info("Count of potential answers has reached 1 or fewer words, breaking revision loop")
			break
		}
	}

	if answerList.Count() == 5237 {
		glog.Warning("After applying logic, answerList is still at maximum count")
		solvingError = "Keep adding letters and colors to generate potential answers..."
		httpStatus = http.StatusBadRequest
		glog.Info("Writing results count")
		var resultCount int = answerList.Count()
		glog.Info("Writing results")
		results := strings.Join(answerList.Words, " ")
		glog.Info("Returning solveWordle function")
		return results, resultCount, solvingError, httpStatus
	}
	
	glog.Info("Writing results count")
	var resultCount int = answerList.Count()
	glog.Info("Writing results")
	results := strings.Join(answerList.Words, " ")
	glog.Info("Returning solveWordle function")
	return results, resultCount, solvingError, httpStatus
}

package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSolveWordle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	webServer := gin.Default()
	webServer.POST("/guesses", parseGuesses)

	t.Run("Test valid input", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "I", Color: "grey"},
			{Character: "R", Color: "grey"},
			{Character: "A", Color: "yellow"},
			{Character: "T", Color: "grey"},
			{Character: "E", Color: "yellow"},
			{Character: "P", Color: "grey"},
			{Character: "L", Color: "grey"},
			{Character: "E", Color: "green"},
			{Character: "A", Color: "green"},
			{Character: "D", Color: "green"},
			{Character: "A", Color: "grey"},
			{Character: "H", Color: "grey"},
			{Character: "E", Color: "green"},
			{Character: "A", Color: "green"},
			{Character: "D", Color: "green"},
			{Character: "S", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "E", Color: "green"},
			{Character: "A", Color: "green"},
			{Character: "D", Color: "green"},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/guesses", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		httpResponse := httptest.NewRecorder()
		webServer.ServeHTTP(httpResponse, req)

		assert.Equal(t, http.StatusOK, httpResponse.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(httpResponse.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		assert.NotEmpty(t, jsonResponse["result"])
		assert.NotEmpty(t, jsonResponse["resultCount"])
		assert.Equal(t, "KNEAD", jsonResponse["result"])
		assert.Equal(t, float64(1), jsonResponse["resultCount"])
	})

	t.Run("Test grey box regex logic", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/guesses", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		httpResponse := httptest.NewRecorder()
		webServer.ServeHTTP(httpResponse, req)

		assert.Equal(t, http.StatusOK, httpResponse.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(httpResponse.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		assert.NotEmpty(t, jsonResponse["result"])
		assert.NotEmpty(t, jsonResponse["resultCount"])
		assert.NotContains(t, jsonResponse["result"], "t")
		assert.Equal(t, float64(3908), jsonResponse["resultCount"])
	})

	t.Run("Test invalid character input", func(t *testing.T) {
		//Define an invalid grid input (invalid color)
		gridData := []CellData{
			{Character: "&", Color: "green"},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/guesses", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		httpResponse := httptest.NewRecorder()
		webServer.ServeHTTP(httpResponse, req)

		assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(httpResponse.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		assert.Contains(t, jsonResponse["error"], "Invalid character: &")
	})

	t.Run("Test invalid color input", func(t *testing.T) {
		//Define an invalid grid input (invalid color)
		gridData := []CellData{
			{Character: "P", Color: "purple"},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/guesses", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		webServer.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		assert.Contains(t, jsonResponse["error"], "Invalid color: purple")
	})
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSolveWordle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/solve", solveWordle)

	t.Run("Test valid input", func(t *testing.T) {
		// Define a valid grid input
		gridData := []BoxData{
        {Character:"i",Color:"grey"},
		{Character:"r",Color:"grey"},
		{Character:"a",Color:"yellow"},
		{Character:"t",Color:"grey"},
		{Character:"e",Color:"yellow"},
		{Character:"p",Color:"grey"},
		{Character:"l",Color:"grey"},
		{Character:"e",Color:"green"},
		{Character:"a",Color:"green"},
		{Character:"d",Color:"green"},
		{Character:"a",Color:"grey"},
		{Character:"h",Color:"grey"},
		{Character:"e",Color:"green"},
		{Character:"a",Color:"green"},
		{Character:"d",Color:"green"},
		{Character:"s",Color:"grey"},
		{Character:"t",Color:"grey"},
		{Character:"e",Color:"green"},
		{Character:"a",Color:"green"},
		{Character:"d",Color:"green"},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
		{Character:"",Color:""},
    }

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/solve", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response["result"])
		assert.NotEmpty(t, response["resultSummary"])
		assert.Equal(t, "knead", response["result"])
		assert.Equal(t, "Potential answers: 1\n", response["resultSummary"])
	})

	t.Run("Test invalid character input", func(t *testing.T) {
		// Define an invalid grid input (invalid color)
		gridData := []BoxData{
			{Character: "&", Color: "green"},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/solve", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Contains(t, response["error"], "Invalid character: &")
	})

	t.Run("Test invalid color input", func(t *testing.T) {
		// Define an invalid grid input (invalid color)
		gridData := []BoxData{
			{Character: "p", Color: "purple"},
		}

		jsonData, _ := json.Marshal(gridData)
		req, _ := http.NewRequest(http.MethodPost, "/solve", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Contains(t, response["error"], "Invalid color: purple")
	})
}

package main

import (
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewAnswerList(t *testing.T) {
	answerList := createNewAnswerList()
	var resultCount int = answerList.Count()

	assert.NotEmpty(t, answerList)
	assert.IsType(t, dictionary_tools.MySimpleDict{}, *answerList)
	assert.Equal(t, 5191, resultCount)
	assert.Contains(t, answerList.Words, "ABACK")
}

func TestSetRegexPatterns(t *testing.T) {

	t.Run("Test simple green logic", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "A", Color: "green"},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
		}
		singleLetterRegexPattern := setSingleLetterRegexPattern(0, "A", gridData)

		assert.NotEmpty(t, singleLetterRegexPattern)
		assert.Equal(t, "A....", singleLetterRegexPattern)
	})

	t.Run("Test simple yellow logic", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "A", Color: "yellow"},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
		}
		singleLetterRegexPattern := setSingleLetterRegexPattern(0, "A", gridData)

		assert.NotEmpty(t, singleLetterRegexPattern)
		assert.Equal(t, "[^A].*[A{1,4}].*", singleLetterRegexPattern)
	})

	t.Run("Test simple grey logic", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "A", Color: "grey"},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
			{Character: "", Color: ""},
		}
		singleLetterRegexPattern := setSingleLetterRegexPattern(0, "A", gridData)

		assert.NotEmpty(t, singleLetterRegexPattern)
		assert.Equal(t, "[^A]....", singleLetterRegexPattern)
	})

	t.Run("Test complex yellow logic #1", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "I", Color: "green"},
			{Character: "R", Color: "grey"},
			{Character: "A", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "E", Color: "grey"},
			{Character: "I", Color: "green"},
			{Character: "G", Color: "grey"},
			{Character: "N", Color: "grey"},
			{Character: "I", Color: "yellow"},
			{Character: "S", Color: "grey"},
		}
		multiLetterRegexPattern := setMultiLetterRegexPattern(8, "I", gridData)

		assert.NotEmpty(t, multiLetterRegexPattern)
		assert.Equal(t, "(?:[^I]*I){2}[^I]*", multiLetterRegexPattern)
	})

	t.Run("Test complex grey logic #1", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "T", Color: "grey"},
			{Character: "A", Color: "yellow"},
			{Character: "R", Color: "green"},
			{Character: "T", Color: "grey"},
			{Character: "S", Color: "yellow"},
		}
		multiLetterRegexPattern := setMultiLetterRegexPattern(3, "T", gridData)

		assert.NotEmpty(t, multiLetterRegexPattern)
		assert.Equal(t, "[^T]*", multiLetterRegexPattern)
	})

	t.Run("Test complex grey logic #2", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "C", Color: "grey"},
			{Character: "A", Color: "yellow"},
			{Character: "R", Color: "green"},
			{Character: "T", Color: "grey"},
			{Character: "S", Color: "yellow"},
		}
		multiLetterRegexPattern := setMultiLetterRegexPattern(3, "T", gridData)

		assert.NotEmpty(t, multiLetterRegexPattern)
		assert.Equal(t, "[^T]*", multiLetterRegexPattern)
	})

	t.Run("Test complex grey logic #3", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "P", Color: "grey"},
			{Character: "U", Color: "green"},
			{Character: "P", Color: "grey"},
			{Character: "P", Color: "green"},
			{Character: "Y", Color: "green"},
		}
		multiLetterRegexPattern := setMultiLetterRegexPattern(0, "P", gridData)
		assert.Empty(t, multiLetterRegexPattern)
		
		multiLetterRegexPattern = setMultiLetterRegexPattern(2, "P", gridData)
		assert.Empty(t, multiLetterRegexPattern)

		multiLetterRegexPattern = setMultiLetterRegexPattern(3, "P", gridData)
		assert.NotEmpty(t, multiLetterRegexPattern)
		assert.Equal(t, "[^P][^P][^P]P[^P]", multiLetterRegexPattern)
	})
}

func TestReviseAnswerList(t *testing.T) {
	answerList := createNewAnswerList()
	regexPattern := "^S....$"

	answerList = reviseAnswerList(answerList, regexPattern)
	var resultCount int = answerList.Count()

	assert.NotEmpty(t, answerList)
	assert.IsType(t, dictionary_tools.MySimpleDict{}, *answerList)
	assert.Equal(t, 683, resultCount)
	assert.Contains(t, answerList.Words, "SABER")
}

func TestNonAlpha(t *testing.T) {

	t.Run("Test letter", func(t *testing.T) {
		nonAlphaTest := nonAlpha("A")
		var exampleBool bool

		assert.Empty(t, nonAlphaTest)
		assert.IsType(t, exampleBool, nonAlphaTest)
		assert.Equal(t, false, nonAlphaTest)
	})
	t.Run("Test number", func(t *testing.T) {
		nonAlphaTest := nonAlpha("3")
		var exampleBool bool

		assert.NotEmpty(t, nonAlphaTest)
		assert.IsType(t, exampleBool, nonAlphaTest)
		assert.Equal(t, true, nonAlphaTest)
	})

	t.Run("Test symbol", func(t *testing.T) {
		nonAlphaTest := nonAlpha("&")
		var exampleBool bool

		assert.NotEmpty(t, nonAlphaTest)
		assert.IsType(t, exampleBool, nonAlphaTest)
		assert.Equal(t, true, nonAlphaTest)
	})

}

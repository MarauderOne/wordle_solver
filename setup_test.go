package main

import (
	"testing"
	"github.com/MarauderOne/wordle_solver/dictionary_tools"
	"github.com/stretchr/testify/assert"
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

	t.Run("Test complex grey logic", func(t *testing.T) {
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
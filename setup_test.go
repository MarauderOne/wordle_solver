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
	assert.Equal(t, 5237, resultCount)
	assert.Contains(t, answerList.Words, "ABACK")
}

func TestSetRegexPatterns(t *testing.T) {

	t.Run("Test simple logic", func(t *testing.T) {
	//Define a valid grid input
	gridData := []CellData{
		{Character: "A", Color: "green"},
		{Character: "", Color: ""},
		{Character: "", Color: ""},
		{Character: "", Color: ""},
		{Character: "", Color: ""},
	}
	greenRegex, yellowRegex, greyRegex := setRegexPatterns(0, "A", gridData)

	assert.NotEmpty(t, greenRegex)
	assert.Equal(t, "A....", greenRegex)
	assert.NotEmpty(t, yellowRegex)
	assert.Equal(t, "[^A][A{1,}]...$|^[^A].[A{1,}]..$|^[^A]..[A{1,}].$|^[^A]...[A{1,}]", yellowRegex)
	assert.NotEmpty(t, greyRegex)
	assert.Equal(t, "[^A]*", greyRegex)
	})

	t.Run("Test complex grey logic", func(t *testing.T) {
		//Define a valid grid input
		gridData := []CellData{
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
			{Character: "T", Color: "grey"},
		}
		greenRegex, yellowRegex, greyRegex := setRegexPatterns(0, "T", gridData)
	
		assert.NotEmpty(t, greenRegex)
		assert.Equal(t, "T....", greenRegex)
		assert.NotEmpty(t, yellowRegex)
		assert.Equal(t, "[^T][T{1,}]...$|^[^T].[T{1,}]..$|^[^T]..[T{1,}].$|^[^T]...[T{1,}]", yellowRegex)
		assert.NotEmpty(t, greyRegex)
		assert.Equal(t, "[^T]....", greyRegex)
		})
}

func TestReviseAnswerList(t *testing.T) {
	answerList := createNewAnswerList()
	regexPattern := "^S....$"

	answerList = reviseAnswerList(answerList, regexPattern)
	var resultCount int = answerList.Count()

	assert.NotEmpty(t, answerList)
	assert.IsType(t, dictionary_tools.MySimpleDict{}, *answerList)
	assert.Equal(t, 689, resultCount)
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
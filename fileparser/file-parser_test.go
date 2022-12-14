package fileparser_test

import (
	"catarinabombaca/advent-code-22/fileparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileParser(t *testing.T) {
	//Given a file path
	filePath := "../inputs/testInput.txt"
	expected := []int{10000, 9000, 5000, 2000}

	//When we extract the file contents
	contents, err := fileparser.ParseToListOfCalories(filePath)
	require.NoError(t, err, "error while parsing the file")

	//Then we get a list of ints
	assert.Equal(t, expected, contents)
	assert.Len(t, contents, 4, "contents are do not have expected length")
}

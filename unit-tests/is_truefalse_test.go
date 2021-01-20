package my_tests

import (
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_True(t *testing.T) {
	// Arrange
	actual := true
	matcher := is.True
	matcherN := is.False

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.EqualTo(true), "Result")
	Assert(t).That(resultN, is.EqualTo(false), "Not Result")
	Assert(t).That(msgN, is.EqualTo(`Expected "false" (bool), but found "true" (bool)`), "Not Message")
}

func Test_False(t *testing.T) {
	// Arrange
	actual := false
	matcher := is.True
	matcherN := is.False

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.EqualTo(false), "Result")
	Assert(t).That(msg, is.EqualTo(`Expected "true" (bool), but found "false" (bool)`), "Message")
	Assert(t).That(resultN, is.EqualTo(true), "Not Result")
}

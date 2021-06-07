package my_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Equal_Equal(t *testing.T) {
	// Arrange
	actual := 0
	expected := 0
	matcher := is.EqualTo(expected)
	matcherN := is.NotEqualTo(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%d" (int)`, actual)), "Not Message")
}

func Test_Equal_NotEqual(t *testing.T) {
	// Arrange
	actual := 0
	expected := 1
	matcher := is.EqualTo(expected)
	matcherN := is.NotEqualTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected "%d" (int), but found "%d" (int)`, expected, actual)), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

func Test_Equal_CustomEquality_Equal(t *testing.T) {
	// Arrange
	actual := 1
	expected := 2
	customEquals := func(act, exp interface{}) bool { return act.(int) == (exp.(int) / 2) }
	matcher := is.EqualToC(expected, customEquals)
	matcherN := is.NotEqualToC(expected, customEquals)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%d" (int)`, expected)), "Not Message")
}

func Test_Equal_CustomEquality_NotEqual(t *testing.T) {
	// Arrange
	actual := 1
	expected := 4
	customEquals := func(act, exp interface{}) bool { return act.(int) == (exp.(int) / 2) }
	matcher := is.EqualToC(expected, customEquals)
	matcherN := is.NotEqualToC(expected, customEquals)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected "%d" (int), but found "%d" (int)`, expected, actual)), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

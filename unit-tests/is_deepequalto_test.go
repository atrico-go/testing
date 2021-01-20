package my_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_DeepEqual_Equal(t *testing.T) {
	// Arrange
	actual := 0
	expected := 0
	matcher := is.DeepEqualTo(expected)
	matcherN := is.NotDeepEqualTo(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%d" (int)`, actual)), "Not Message")
}

func Test_DeepEqual_NotEqual(t *testing.T) {
	// Arrange
	actual := 0
	expected := 1
	matcher := is.DeepEqualTo(expected)
	matcherN := is.NotDeepEqualTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected "%d" (int), but found "%d" (int)`, expected, actual)), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

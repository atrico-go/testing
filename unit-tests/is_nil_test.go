package my_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Nil_Nil(t *testing.T) {
	// Arrange
	var actual interface{} = nil
	matcher := is.Nil
	matcherN := is.NotNil

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(`Expected something other than "nil"`), "Not Message")
}

func Test_Nil_NotNil(t *testing.T) {
	// Arrange
	actual := TestType{}
	matcher := is.Nil
	matcherN := is.NotNil

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected "nil", but found "%v" (%T)`, actual, actual)), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

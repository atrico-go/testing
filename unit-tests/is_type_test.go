package my_tests

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

type TestType2 struct{}

func Test_Type_SameType(t *testing.T) {
	// Arrange
	actual := TestType{}
	expected := reflect.TypeOf(TestType{})
	matcher := is.Type(expected)
	matcherN := is.NotType(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected a type other than "%v"`, expected)), "Not Message")
}

func Test_Type_DifferentType(t *testing.T) {
	// Arrange
	actual := TestType2{}
	expected := reflect.TypeOf(TestType{})
	matcher := is.Type(expected)
	matcherN := is.NotType(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected type "%v", but found "%T" (%v)`, expected, actual, actual)), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

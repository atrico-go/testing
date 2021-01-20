package my_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Equivalent_Equal(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3}
	expected := []int{1, 2, 3}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%v"`, actual)), "Not Message")

}

func Test_Equivalent_WrongOrder(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3}
	expected := []int{2, 1, 3}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%v"`, expected)), "Not Message")
}

func Test_Equivalent_Duplicate(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3, 1}
	expected := []int{2, 1, 3, 1}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, _ := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	Assert(t).That(result, is.True, "Result")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(fmt.Sprintf(`Expected something other than "%v"`, expected)), "Not Message")
}

func Test_Equivalent_ExtraItems(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected (equivalent to) "%d", but found "%d", Extra items "%v"`, expected, actual, []int{4, 5})), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

func Test_Equivalent_MissingItems(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4, 5}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected (equivalent to) "%d", but found "%d", Missing items "%v"`, expected, actual, []int{4, 5})), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

func Test_Equivalent_ExtraAndMissingItems(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 6, 7}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected (equivalent to) "%d", but found "%d", Extra items "%v", Missing items "%v"`, expected, actual, []int{4, 5}, []int{6, 7})), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

func Test_Equivalent_DuplicateExtraAndMissingItems(t *testing.T) {
	// Arrange
	actual := []int{1, 2, 3, 4, 4}
	expected := []int{1, 2, 3, 3, 4}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, _ := matcherN(actual)

	// Assert
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(fmt.Sprintf(`Expected (equivalent to) "%d", but found "%d", Extra items "%v", Missing items "%v"`, expected, actual, []int{4}, []int{3})), "Message")
	Assert(t).That(resultN, is.True, "Not Result")
}

func Test_Equivalent_NotASlice(t *testing.T) {
	// Arrange
	actual := struct{}{}
	expected := []int{1, 2, 3}
	matcher := is.EquivalentTo(expected)
	matcherN := is.NotEquivalentTo(expected)

	// Act
	result, msg := matcher(actual)
	resultN, msgN := matcherN(actual)

	// Assert
	expMsg := fmt.Sprintf(`Value is not a slice: "%v" (%T)`, actual, actual)
	Assert(t).That(result, is.False, "Result")
	Assert(t).That(msg, is.EqualTo(expMsg), "Message")
	Assert(t).That(resultN, is.False, "Not Result")
	Assert(t).That(msgN, is.EqualTo(expMsg), "Not Message")
}

func Test_Equivalent_ExpectedNotASlice(t *testing.T) {
	// Arrange
	expected := 123

	// Act
	thePanic := PanicCatcher(func() { is.EquivalentTo(expected) })
	thePanicN := PanicCatcher(func() { is.NotEquivalentTo(expected) })

	// Assert
	expectedPanic := fmt.Sprintf(`Expected value is not a slice: "%v" (%T)`, expected, expected)
	Assert(t).That(thePanic, is.EqualTo(expectedPanic), "Panic")
	Assert(t).That(thePanicN, is.EqualTo(expectedPanic), "Not Panic")
}

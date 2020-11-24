package is

import (
	"reflect"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

func DeepEqualTo(expected interface{}) Matcher {
	return CreateMatcher(deepEqualsMatch(expected), messages.ExpectedButActual(expected))
}

func NotDeepEqualTo(expected interface{}) Matcher {
	return CreateNotMatcher(deepEqualsMatch(expected), messages.ExpectedOtherThan(expected))
}

func deepEqualsMatch(expected interface{}) MatcherImplementation {
	return func(actual interface{}) bool { return reflect.DeepEqual(actual, expected) }
}

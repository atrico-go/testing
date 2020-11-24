package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

func EqualTo(expected interface{}) Matcher {
	return CreateMatcher(equalsMatch(expected), messages.ExpectedButActual(expected))
}

func NotEqualTo(expected interface{}) Matcher {
	return CreateNotMatcher(equalsMatch(expected), messages.ExpectedOtherThan(expected))
}

func equalsMatch(expected interface{}) MatcherImplementation {
	return func(actual interface{}) bool { return actual == expected }
}

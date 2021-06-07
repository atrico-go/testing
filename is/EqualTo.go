package is

import (
	"reflect"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

type CustomEquality func(act, exp interface{}) bool

func EqualTo(expected interface{}) Matcher {
	return EqualToC(expected, defaultEquality)
}

func NotEqualTo(expected interface{}) Matcher {
	return NotEqualToC(expected, defaultEquality)
}

func EqualToC(expected interface{}, equality CustomEquality) Matcher {
	return CreateMatcher(equalsMatchC(expected, equality), messages.ExpectedButActual(expected))
}

func NotEqualToC(expected interface{}, equality CustomEquality) Matcher {
	return CreateNotMatcher(equalsMatchC(expected, equality), messages.ExpectedOtherThan(expected))
}

func equalsMatch(expected interface{}) MatcherImplementation {
	return equalsMatchC(expected, defaultEquality)
}

func equalsMatchC(expected interface{}, equality CustomEquality) MatcherImplementation {
	return func(actual interface{}) bool { return equality(actual, expected) }
}

var defaultEquality = func(act, exp interface{}) bool { return reflect.DeepEqual(act, exp) }

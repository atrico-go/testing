package is

import (
	"reflect"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

func Type(expected reflect.Type) Matcher {
	return CreateMatcher(typeMatch(expected), messages.ExpectedTypeButActual(expected))
}

func NotType(expected reflect.Type) Matcher {
	return CreateNotMatcher(typeMatch(expected), messages.ExpectedTypeButActual(expected))
}

func typeMatch(expected reflect.Type) MatcherImplementation {
	return func(actual interface{}) bool { return reflect.TypeOf(actual) == expected }
}

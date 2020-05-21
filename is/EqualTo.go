package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

func EqualTo(expected interface{}) Matcher {
	return func(actual interface{}) (bool, string) {
		if actual == expected {
			return true, ""
		}
		return false, messages.ExpectedButActual(expected, actual)
	}
}

func NotEqualTo(expected interface{}) Matcher {
	return func(actual interface{}) (bool, string) {
		if actual != expected {
			return true, ""
		}
		return false, messages.ExpectedOtherThan(expected)
	}
}

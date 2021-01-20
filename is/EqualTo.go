package is

import (
	"fmt"
	. "github.com/atrico-go/testing/assert"
)

func EqualTo(expected interface{}) Matcher {
	return func(actual interface{}) (bool, string) {
		if actual == expected {
			return true, ""
		}
		return false, fmt.Sprintf("Expected %v, but found %v", expected, actual)
	}
}

func NotEqualTo(expected interface{}) Matcher {
	return func(actual interface{}) (bool, string) {
		if actual != expected {
			return true, ""
		}
		return false, fmt.Sprintf("Expected something other than %v", expected)
	}
}

package is

import (
	"atrico.net/go/assert"
	"fmt"
)

func EqualTo(expected interface{}) assert.Matcher {
	return equalTo(expected)
}

func NotEqualTo(expected interface{}) assert.Matcher {
	return equalTo(expected).NotModifier()
}

func equalTo(expected interface{}) assert.Matcher {
	return func(actual interface{}) (bool, string) {
		if actual == expected {
			return true, fmt.Sprintf("%v == %v", actual, expected)
		}
		return false, fmt.Sprintf("%v != %v", actual, expected)
	}
}

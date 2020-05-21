package is

import (
	"fmt"
	. "github.com/atrico-go/testing/assert"
)

var Nil Matcher = func(actual interface{}) (bool, string) {
	if actual == nil {
		return true, ""
	}
	return false, fmt.Sprintf("Expected nil, but found %v", actual)
}

var NotNil Matcher = func(actual interface{}) (bool, string) {
	if actual != nil {
		return true, ""
	}
	return false, "Expected something other than nil"
}

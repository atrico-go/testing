package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

var True Matcher = func(actual interface{}) (bool, string) {
	if actual == true {
		return true, ""
	}
	return false, messages.ExpectedButActual("true", actual)
}

var False Matcher = func(actual interface{}) (bool, string) {
	if actual == false {
		return true, ""
	}
	return false, messages.ExpectedButActual("false", actual)
}

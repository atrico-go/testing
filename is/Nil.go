package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

var Nil Matcher = func(actual interface{}) (bool, string) {
	if actual == nil {
		return true, ""
	}
	return false, 		messages.ExpectedButActual("nil", actual)

}

var NotNil Matcher = func(actual interface{}) (bool, string) {
	if actual != nil {
		return true, ""
	}
	return false, messages.ExpectedOtherThan("nil")
}

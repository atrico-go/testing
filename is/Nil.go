package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

var Nil = CreateMatcher(nilMatch, messages.ExpectedButActual(nil))

var NotNil = CreateNotMatcher(nilMatch, messages.ExpectedOtherThan(nil))

var nilMatch = equalsMatch(nil)

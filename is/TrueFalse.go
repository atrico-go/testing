package is

import (
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

var True = CreateMatcher(equalsMatch(true), messages.ExpectedButActual(true))

var False = CreateMatcher(equalsMatch(false), messages.ExpectedButActual(false))

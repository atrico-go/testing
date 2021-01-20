package messages

import (
	"fmt"
	"reflect"

	"github.com/atrico-go/testing/assert"
)

func ExpectedButActual(expected interface{}) assert.MessageProvider {
	if expected == nil {
		return func(actual interface{}) string { return fmt.Sprintf(`Expected "nil", but found "%v" (%T)`, actual, actual) }
	}
	return func(actual interface{}) string { return fmt.Sprintf(`Expected "%v" (%T), but found "%v" (%T)`, expected, expected, actual, actual) }
}

func ExpectedTypeButActual(expected reflect.Type) assert.MessageProvider {
	return func(actual interface{}) string { return fmt.Sprintf(`Expected type "%v", but found "%T" (%v)`, expected, actual, actual) }
}

func ExpectedOtherThan(expected interface{}) assert.MessageProvider {
	if expected == nil {
		return func(actual interface{}) string { return `Expected something other than "nil"` }
	}
	return func(actual interface{}) string { return fmt.Sprintf(`Expected something other than "%v" (%T)`, expected, expected) }
}

func ExpectedTypeOtherThan(expected reflect.Type) assert.MessageProvider {
	return func(actual interface{}) string { return fmt.Sprintf(`Expected a type other than "%v"`, expected) }
}

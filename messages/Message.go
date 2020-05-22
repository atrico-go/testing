package messages

import (
	"fmt"
	"github.com/atrico-go/testing/assert"
	"reflect"
)

func ExpectedButActual(expected interface{}) assert.MessageProvider {
	return func (actual interface{}) string {return fmt.Sprintf("Expected %v, but found %v", expected, actual)}
}

func ExpectedTypeButActual(expected reflect.Type) assert.MessageProvider {
	return func (actual interface{}) string {return fmt.Sprintf("Expected type %v, but found %v (%v)", expected, reflect.TypeOf(actual), actual)}
}

func ExpectedOtherThan(expected interface{})  assert.MessageProvider {
	return func (actual interface{}) string {return fmt.Sprintf("Expected something other than %v", expected)}
}

func ExpectedTypeOtherThan(expected reflect.Type, actual interface{}) string {
	return fmt.Sprintf("Expected a type other than %v but found %v", expected, actual)
}
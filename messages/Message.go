package messages

import "fmt"

func ExpectedButActual(expected interface{},actual interface{}) string {
	return fmt.Sprintf("Expected %v, but found %v", expected, actual)
}

func ExpectedOtherThan(actual interface{}) string {
	return fmt.Sprintf( "Expected something other than %v",  actual)
}
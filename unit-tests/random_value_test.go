package my_tests

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/atrico-go/testing/random"
)

func Test_RandomValue(t *testing.T) {
	// Pod types
	t.Run("string", randomValueImpl(new(string)))
	t.Run("rune", randomValueImpl(new(rune)))
	t.Run("int", randomValueImpl(new(int)))
	t.Run("int8", randomValueImpl(new(int8)))
	t.Run("int16", randomValueImpl(new(int16)))
	t.Run("int32", randomValueImpl(new(int32)))
	t.Run("int64", randomValueImpl(new(int64)))
	t.Run("uint", randomValueImpl(new(uint)))
	t.Run("uint8", randomValueImpl(new(uint8)))
	t.Run("uint16", randomValueImpl(new(uint16)))
	t.Run("uint32", randomValueImpl(new(uint32)))
	t.Run("uint64", randomValueImpl(new(uint64)))
	t.Run("float32", randomValueImpl(new(float64)))
	t.Run("float64", randomValueImpl(new(float64)))
	t.Run("bool", randomValueImpl(new(bool)))
	// Pointers
	t.Run("*string", randomValueImpl(new(*string)))
	t.Run("*rune", randomValueImpl(new(*rune)))
	t.Run("*int", randomValueImpl(new(*int)))
	t.Run("*int8", randomValueImpl(new(*int8)))
	t.Run("*int16", randomValueImpl(new(*int16)))
	t.Run("*int32", randomValueImpl(new(*int32)))
	t.Run("*int64", randomValueImpl(new(*int64)))
	t.Run("*uint", randomValueImpl(new(*uint)))
	t.Run("*uint8", randomValueImpl(new(*uint8)))
	t.Run("*uint16", randomValueImpl(new(*uint16)))
	t.Run("*uint32", randomValueImpl(new(*uint32)))
	t.Run("*uint64", randomValueImpl(new(*uint64)))
	t.Run("*float32", randomValueImpl(new(*float64)))
	t.Run("*float64", randomValueImpl(new(*float64)))
	t.Run("*bool", randomValueImpl(new(*bool)))
	// Slice types
	t.Run("[]string", randomValueImpl(new([]string)))
	t.Run("[]rune", randomValueImpl(new([]rune)))
	t.Run("[]int", randomValueImpl(new([]int)))
	t.Run("[]int8", randomValueImpl(new([]int8)))
	t.Run("[]int16", randomValueImpl(new([]int16)))
	t.Run("[]int32", randomValueImpl(new([]int32)))
	t.Run("[]int64", randomValueImpl(new([]int64)))
	t.Run("[]uint", randomValueImpl(new([]uint)))
	t.Run("[]uint8", randomValueImpl(new([]uint8)))
	t.Run("[]uint16", randomValueImpl(new([]uint16)))
	t.Run("[]uint32", randomValueImpl(new([]uint32)))
	t.Run("[]uint64", randomValueImpl(new([]uint64)))
	t.Run("[]float32", randomValueImpl(new([]float64)))
	t.Run("[]float64", randomValueImpl(new([]float64)))
	t.Run("[]bool", randomValueImpl(new([]bool)))
	// Map
	t.Run("map[string]int", randomValueImpl(new(map[string]int)))
	// Struct
	t.Run("struct", randomValueImpl(new(testStruct)))
	t.Run("[]struct", randomValueImpl(new([]testStruct)))
}

type testStruct struct {
	One string
	Two int
	Three []string
}

func randomValueImpl(receiver interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		// Arrange
		rn := random.NewValueGenerator()

		// Act
		err := rn.Value(receiver)
		fmt.Printf("Value = '%v' (%v)\n", reflect.ValueOf(receiver).Elem(), reflect.TypeOf(receiver).Elem())
		if reflect.TypeOf(receiver).Elem().Kind() == reflect.Ptr {
			fmt.Printf("Points to = '%v'\n", reflect.ValueOf(receiver).Elem().Elem())
		}
		// Assert
		Assert(t).That(err, is.Nil, "No error")
	}
}

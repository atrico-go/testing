package my_tests

import (
	. "github.com/atrico-go/testing/assert"
	. "github.com/atrico-go/testing/is"
	"reflect"
	"testing"
)

type TestType struct{}
type TestType2 struct{}

func TestEqual_pass(t *testing.T) {
	actual := 0
	expected := 0
	Assert(t).That(actual, EqualTo(expected), "")
}
func TestEqual_fail(t *testing.T) {
	actual := 0
	expected := 1
	Assert(t).That(actual, EqualTo(expected), "Reason")
}

func TestNotEqual_pass(t *testing.T) {
	actual := 0
	expected := 1
	Assert(t).That(actual, NotEqualTo(expected), "")
}
func TestNotEqual_fail(t *testing.T) {
	actual := 0
	expected := 0
	Assert(t).That(actual, NotEqualTo(expected), "")
}

func TestNil_pass(t *testing.T) {
	var actual interface{} = nil
	Assert(t).That(actual, Nil, "")
}
func TestNil_fail(t *testing.T) {
	actual := TestType{}
	Assert(t).That(actual, Nil, "Reason")
}

func TestNotNil_pass(t *testing.T) {
	actual := TestType{}
	Assert(t).That(actual, NotNil, "")
}
func TestNotNil_fail(t *testing.T) {
	var actual interface{} = nil
	Assert(t).That(actual, NotNil, "Reason")
}

func TestType_pass(t *testing.T) {
	actual := TestType{}
	expected := reflect.TypeOf(TestType{})
	Assert(t).That(actual, Type(expected), "")
}
func TestType_fail(t *testing.T) {
	actual := TestType2{}
	expected := reflect.TypeOf(TestType{})
	Assert(t).That(actual, Type(expected), "Reason")
}

func TestNotType_pass(t *testing.T) {
	actual := TestType2{}
	expected := reflect.TypeOf(TestType{})
	Assert(t).That(actual, NotType(expected), "")
}
func TestNotType_fail(t *testing.T) {
	actual := TestType{}
	expected := reflect.TypeOf(TestType{})
	Assert(t).That(actual, NotType(expected), "Reason")
}

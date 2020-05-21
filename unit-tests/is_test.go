package my_tests

import (
	. "github.com/atrico-go/testing/assert"
	. "github.com/atrico-go/testing/is"
	"testing"
)

func TestEqual_pass(t *testing.T) {
	actual := 0
	expected := 0
	Assert(t).That(actual, EqualTo(expected))
}
func TestEqual_fail(t *testing.T) {
	actual := 0
	expected := 1
	Assert(t).That(actual, EqualTo(expected))
}

func TestNotEqual_pass(t *testing.T) {
	actual := 0
	expected := 1
	Assert(t).That(actual, NotEqualTo(expected))
}
func TestNotEqual_fail(t *testing.T) {
	actual := 0
	expected := 0
	Assert(t).That(actual, NotEqualTo(expected))
}

type TestObj struct {}

func TestNil_pass(t *testing.T) {
	var actual interface{} = nil
	Assert(t).That(actual, Nil)
}
func TestNil_fail(t *testing.T) {
	actual := TestObj{}
	Assert(t).That(actual, Nil)
}

func TestNotNil_pass(t *testing.T) {
	actual := TestObj{}
	Assert(t).That(actual, NotNil)
}
func TestNotNil_fail(t *testing.T) {
	var actual interface{} = nil
	Assert(t).That(actual, NotNil)
}
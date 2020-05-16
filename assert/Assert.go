package assert

import "testing"

// Wrapper around T to allow extension functions
type Assert struct {
	T *testing.T
}

func (assert Assert) Fail(format string, args ...interface{}) {
	assert.T.Errorf(format, args...)
	assert.T.Fail()
}

func (assert Assert) That(actual interface{}, matcher Matcher) {
	pass,msg := matcher(actual)
	if !pass{
		assert.Fail(msg)
	}
}

type Matcher func(interface{}) (pass bool,message string)

func (matcher Matcher) NotModifier() Matcher {
	return func(actual interface{}) (pass bool,message string) {
		pass,msg := matcher(actual)
		return !pass, msg
	}
}
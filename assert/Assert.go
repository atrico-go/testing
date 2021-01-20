package assert

import "testing"

func Assert(t *testing.T) AssertWrapper {
	return AssertWrapper{t}
}

// Wrapper around T to allow extension functions
type AssertWrapper struct {
	T *testing.T
}

func (assert AssertWrapper) Fail(format string, args ...interface{}) {
	assert.T.Errorf(format, args...)
	assert.T.Fail()
}

func (assert AssertWrapper) That(actual interface{}, matcher Matcher) {
	pass, message := matcher(actual)
	if !pass {
		assert.Fail(message)
	}
}

type Matcher func(interface{}) (bool, string)

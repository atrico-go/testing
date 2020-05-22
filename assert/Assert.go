package assert

import (
	"testing"
)

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

type Matcher func(actual interface{}) (bool, string)
type MatcherImplementation func(actual interface{}) bool
type MessageProvider func(actual interface{}) string

// Helper functions
func CreateMatcher(match MatcherImplementation, message MessageProvider) Matcher {
	return func(actual interface{}) (bool, string) {
		if match(actual) {
			return true, ""
		}
		return false, message(actual)
	}
}
func CreateNotMatcher(match MatcherImplementation, message MessageProvider) Matcher {
	return CreateMatcher(func (actual interface{}) bool {return !match(actual)}, message)
}
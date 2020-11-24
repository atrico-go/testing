package assert

import (
	"fmt"
	"strings"
	"testing"
)

func Assert(t *testing.T) Wrapper {
	return Wrapper{t}
}

// Wrapper around T to allow extension functions
type Wrapper struct {
	T *testing.T
}

func (assert Wrapper) Fail(format string, args ...interface{}) {
	assert.Logf(format, args...)
}

func (assert Wrapper) Logf(format string, args ...interface{}) {
	assert.T.Logf(format, args...)
	assert.T.FailNow()
}

func (assert Wrapper) That(actual interface{}, matcher Matcher, reasonFormat string, reasonArgs ...interface{}) {
	pass, message := matcher(actual)
	if !pass {
		output := fmt.Sprintf(reasonFormat, reasonArgs...)
		sep := ""
		if len(output) > 0 {
			sep = " => "
		}
		assert.Fail(strings.Join([]string{output, message}, sep))
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
	return CreateMatcher(func(actual interface{}) bool { return !match(actual) }, message)
}

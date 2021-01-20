package random

import (
	"math/rand"
)

type ValueGenerator interface {
	String() string
	StringOfLen(length int) string
	Int() int
	IntUpto(max int) int
	IntBetween(min, max int) int
	Bool() bool
}

func NewValueGenerator() ValueGenerator {
	return newDefaultGenerator()
}

// ----------------------------------------------------------------------------------------------------------------------------
// Builder
// ----------------------------------------------------------------------------------------------------------------------------
type ValueGeneratorBuilder interface {
	WithDefaultStringLength(length int) ValueGeneratorBuilder
	Build() ValueGenerator
}

func NewValueGeneratorBuilder() ValueGeneratorBuilder {
	gen := newDefaultGenerator()
	return &gen
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
var defaultStringLength = 5

type randomValueGenerator struct {
	defaultStringLength int
}

func newDefaultGenerator() randomValueGenerator {
	return randomValueGenerator{defaultStringLength}
}

func (r randomValueGenerator) String() string {
	return r.StringOfLen(r.defaultStringLength)
}

func (r randomValueGenerator) StringOfLen(length int) string {
	text := make([]byte, length)
	for i := range text {
		n := rand.Intn(62)
		switch {
		case 0 <= n && n < 10:
			text[i] = byte('0' + n)
		case 10 <= n && n < 36:
			text[i] = byte('A' + n - 10)
		default:
			text[i] = byte('a' + n - 36)
		}
	}
	return string(text)
}

func (r randomValueGenerator) Int() int {
	return rand.Int()
}

func (r randomValueGenerator) IntUpto(max int) int {
	return rand.Intn(max)
}

func (r randomValueGenerator) IntBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

func (r randomValueGenerator) Bool() bool {
	return rand.Intn(2) == 0
}

func (r *randomValueGenerator) WithDefaultStringLength(value int) ValueGeneratorBuilder {
	r.defaultStringLength = value
	return r
}

func (r *randomValueGenerator) Build() ValueGenerator {
	return *r
}

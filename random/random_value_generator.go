package random

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type ValueGenerator interface {
	String() string
	StringOfLen(length int) string
	Rune() rune
	Int() int
	IntUpto(max int) int
	IntBetween(min, max int) int
	Uint() uint
	Float32() float32
	Float64() float64
	Bool() bool

	// Get a value of any (supported) type
	Value(receiver interface{}) error
}

func NewValueGenerator() ValueGenerator {
	return newDefaultGenerator()
}

// ----------------------------------------------------------------------------------------------------------------------------
// Builder
// ----------------------------------------------------------------------------------------------------------------------------
type ValueGeneratorBuilder interface {
	WithDefaultStringLength(length int) ValueGeneratorBuilder
	WithDefaultSliceLength(length int) ValueGeneratorBuilder
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
var defaultSliceLength = 3

type randomValueGenerator struct {
	rn                  *rand.Rand
	defaultStringLength int
	defaultSliceLength  int
}

func newDefaultGenerator() randomValueGenerator {
	return randomValueGenerator{rand.New(rand.NewSource(time.Now().UnixNano())), defaultStringLength, defaultSliceLength}
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
	return r.rn.Int()
}

func (r randomValueGenerator) Rune() rune {
	return rune(r.rn.Uint32())
}

func (r randomValueGenerator) IntUpto(max int) int {
	return r.rn.Intn(max)
}

func (r randomValueGenerator) IntBetween(min, max int) int {
	return r.rn.Intn(max-min) + min
}

func (r randomValueGenerator) Uint() uint {
	return uint(r.rn.Uint32())
}

func (r randomValueGenerator) Float32() float32 {
	return r.rn.Float32()
}

func (r randomValueGenerator) Float64() float64 {
	return r.rn.Float64()
}

func (r randomValueGenerator) Bool() bool {
	return r.rn.Intn(2) == 0
}

func (r randomValueGenerator) Value(receiver interface{}) (err error) {
	rValue := reflect.ValueOf(receiver)
	// Must be a pointer
	if rValue.Kind() == reflect.Ptr {
		var val interface{}
		if val, err = r.createValue(rValue.Type().Elem()); err == nil {
			rValue.Elem().Set(reflect.ValueOf(val))
		}
	} else {
		err = errors.New("receiver is not a pointer")
	}
	return err
}

func (r randomValueGenerator) createValue(rType reflect.Type) (value interface{}, err error) {
	var val interface{}
	// Check type
	switch rType.Kind() {
	case reflect.String:
		value = r.String()
	case reflect.Int:
		value = r.Int()
	case reflect.Int8:
		value = int8(r.Int())
	case reflect.Int16:
		value = int16(r.Int())
	case reflect.Int32:
		value = int32(r.Int())
	case reflect.Int64:
		value = (int64(r.Int()) << 32) | int64(r.Int())
	case reflect.Uint:
		value = r.Uint()
	case reflect.Uint8:
		value = uint8(r.Uint())
	case reflect.Uint16:
		value = uint16(r.Uint())
	case reflect.Uint32:
		value = r.rn.Uint32()
	case reflect.Uint64:
		value = r.rn.Uint64()
	case reflect.Float32:
		value = r.Float32()
	case reflect.Float64:
		value = r.Float64()
	case reflect.Bool:
		value = r.Bool()
	case reflect.Ptr:
		vVal := reflect.New(rType.Elem())
		value = vVal.Interface()
		if val, err = r.createValue(rType.Elem()); err == nil {
			vVal.Elem().Set(reflect.ValueOf(val))
		}
	case reflect.Slice:
		vVal := reflect.MakeSlice(rType, r.defaultSliceLength, r.defaultSliceLength)
		value = vVal.Interface()
		eType := rType.Elem()
		for i := 0; i < r.defaultSliceLength; i++ {
			if val, err = r.createValue(eType); err == nil {
				vVal.Index(i).Set(reflect.ValueOf(val))
			}
		}
	case reflect.Map:
		vVal := reflect.MakeMap(rType)
		value = vVal.Interface()
		tKey := rType.Key()
		tElem := rType.Elem()
		vVal.SetMapIndex(reflect.ValueOf("str"), reflect.ValueOf(123))
		for i := 0; i < r.defaultSliceLength; i++ {
			if kVal, kErr := r.createValue(tKey); kErr == nil {
				if mVal, mErr := r.createValue(tElem); mErr == nil {
					vVal.SetMapIndex(reflect.ValueOf(kVal), reflect.ValueOf(mVal))
				} else {
					return value, mErr
				}
			} else {
				return value, kErr
			}
		}
	case reflect.Struct:
		vVal := reflect.New(rType).Elem()
		value = vVal.Interface()
		for i := 0; i < vVal.NumField(); i++ {
			vField := vVal.Field(i)
			if val, err = r.createValue(vField.Type()); err == nil {
				vField.Set(reflect.ValueOf(val))
			}
		}
	default:
		err = errors.New(fmt.Sprintf("unsupported type: %v", rType.Kind()))
	}
	return value, err
}

func (r *randomValueGenerator) WithDefaultStringLength(value int) ValueGeneratorBuilder {
	r.defaultStringLength = value
	return r
}

func (r *randomValueGenerator) WithDefaultSliceLength(value int) ValueGeneratorBuilder {
	r.defaultSliceLength = value
	return r
}

func (r *randomValueGenerator) Build() ValueGenerator {
	return *r
}

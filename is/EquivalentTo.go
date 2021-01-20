package is

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/messages"
)

func EquivalentTo(expected interface{}) Matcher {
	exp := getExpectedSlice(expected)
	return func(actual interface{}) (bool, string) {
		if act, errMsg := getActualSlice(actual); act == nil {
			return false, errMsg
		} else {
			return equivalentToMatcher(act, exp)
		}
	}
}

func NotEquivalentTo(expected interface{}) Matcher {
	exp := getExpectedSlice(expected)
	return func(actual interface{}) (bool, string) {
		if act, errMsg := getActualSlice(actual); act == nil {
			return false, errMsg

		} else {
			ok, _ := equivalentToMatcher(act, exp)
			if !ok {
				return true, ""
			}
			return false, messages.ExpectedOtherThanNoType(exp)(act)
		}
	}
}

func getExpectedSlice(expected interface{}) []interface{} {
	if slc, ok := takeSliceArg(expected); ok {
		return slc
	}
	panic(fmt.Sprintf(`Expected value is not a slice: "%v" (%T)`, expected, expected))
}

func getActualSlice(actual interface{}) ([]interface{}, string) {
	if slc, ok := takeSliceArg(actual); ok {
		return slc, ""
	}
	return nil, fmt.Sprintf(`Value is not a slice: "%v" (%T)`, actual, actual)
}

func equivalentToMatcher(actual, expected []interface{}) (bool, string) {
	extra := makeIndexSet(actual)
	missing := makeIndexSet(expected)
	for actIdx, actItem := range actual {
		for expIdx := range missing {
			if reflect.DeepEqual(actItem, expected[expIdx]) {
				delete(extra, actIdx)
				delete(missing, expIdx)
				break
			}
		}
	}
	// Success
	if len(extra) == 0 && len(missing) == 0 {
		return true, ""
	}
	msg := strings.Builder{}
	msg.WriteString(fmt.Sprintf(`Expected (equivalent to) "%v", but found "%v"`, expected, actual))
	if len(extra) > 0 {
		msg.WriteString(fmt.Sprintf(`, Extra items "%v"`, mapIndicesToValues(extra, actual)))
	}
	if len(missing) > 0 {
		msg.WriteString(fmt.Sprintf(`, Missing items "%v"`, mapIndicesToValues(missing, expected)))
	}
	return false, msg.String()

}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type intSet map[int]interface{}

func makeIndexSet(slc []interface{}) intSet {
	out := make(intSet, len(slc))
	for idx := range slc {
		out[idx] = nil
	}
	return out
}

func mapIndicesToValues(indices intSet, slc []interface{}) []interface{} {
	idxSlice := make([]int, len(indices))
	i := 0
	for k := range indices {
		idxSlice[i] = k
		i++
	}
	sort.Slice(idxSlice, func(i, j int) bool { return idxSlice[i] < idxSlice[j] })
	out := make([]interface{}, len(idxSlice))
	for i, v := range idxSlice {
		out[i] = slc[v]
	}
	return out
}

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

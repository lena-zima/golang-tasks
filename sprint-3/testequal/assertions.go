//go:build !solution

package testequal

import(
	"reflect"
)

func IsEqual(expected, actual interface{}) bool{
	/*if reflect.ValueOf(expected).Type() != reflect.ValueOf(actual).Type(){
		return false
	}*/

	switch expected.(type) {
	case []int, []byte, map[string]string, string:
		return reflect.DeepEqual(expected, actual)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return expected == actual
	default:
		return false
	}
}
	

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	if IsEqual(expected,actual) {
		return true
	}

	if len(msgAndArgs) != 0{
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}else {
		t.Errorf("")
	}

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	if !IsEqual(expected,actual) {
		return true
	}

	if len(msgAndArgs) != 0{
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}else {
		t.Errorf("")
	}

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {

	t.Helper()

	if IsEqual(expected,actual) {
		return
	}

	if len(msgAndArgs) != 0{
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}else {
		t.Errorf("")
	}

	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()

	if !IsEqual(expected,actual) {
		return
	}

	if len(msgAndArgs) != 0{
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}else {
		t.Errorf("")
	}

	t.FailNow()
}
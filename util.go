package behaviortree

import (
	"reflect"
	"runtime"
	"strings"
)

// getFuncName returns the name of the passed in function
func getFuncName(f interface{}) string {
	fnm := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(fnm, ".")[1]
}

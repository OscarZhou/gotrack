package track

import (
	"os"
	"reflect"
	"runtime"
	"strings"
)

func findFuncName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return name[strings.LastIndex(name, string(os.PathSeparator))+1:]
}

package track

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

type ThreadID int

// Track records all track information
type Track struct {
	Debug bool
	mutex *sync.Mutex
	idx   int
	fn    map[ThreadID]time.Time
}

//
func Default() *Track {
	return &Track{
		Debug: true,
		mutex: &sync.Mutex{},
		idx:   0,
	}
}

//
func (t *Track) Start() {
	if t.Debug {
		t.print()
	}
}

//
func (t *Track) End() {
	if t.Debug {
		t.print()
	}
}

//
func (t *Track) print() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	fpcs := make([]uintptr, 2)
	_ = runtime.Callers(2, fpcs)
	caller := runtime.FuncForPC(fpcs[0])
	actionName := caller.Name()[strings.LastIndex(caller.Name(), ".")+1:]
	caller = runtime.FuncForPC(fpcs[1])
	fmt.Printf("%s function:\t%s \n", actionName, caller.Name())
}

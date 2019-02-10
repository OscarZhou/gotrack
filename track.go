package track

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// CallerName representing the caller function name
type CallerName string

const (
	// PhaseStart stands for function start phase
	PhaseStart = "Start"
	// PhaseEnd stands for function end phase
	PhaseEnd = "End"
)

// Track records all track information
type Track struct {
	Debug   bool
	mutex   *sync.Mutex
	callers map[CallerName]time.Time
}

//
func Default() *Track {
	return &Track{
		Debug:   true,
		mutex:   &sync.Mutex{},
		callers: make(map[CallerName]time.Time),
	}
}

//
func (t *Track) Start() {
	if t.Debug {
		curCaller := t.callerName()
		t.callers[curCaller] = time.Now()
		t.print(PhaseStart, curCaller, 0)
	}
}

//
func (t *Track) End() {
	if t.Debug {
		curCaller := t.callerName()
		elapsed := time.Since(t.callers[curCaller])
		t.print(PhaseEnd, curCaller, elapsed)
		delete(t.callers, curCaller)
	}
}

//
func (t *Track) print(p string, s CallerName, elapsed time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	switch p {
	case PhaseStart:
		fmt.Printf("%s function:\t%v \n", p, string(s))
	case PhaseEnd:
		fmt.Printf("%s function:\t%v took %v \n", p, string(s), elapsed)
	}
}

//
func (t *Track) callerName() CallerName {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	fpcs := make([]uintptr, 1)
	_ = runtime.Callers(3, fpcs)
	caller := runtime.FuncForPC(fpcs[0])
	return CallerName(caller.Name())
}

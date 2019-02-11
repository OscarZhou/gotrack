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
	// PhaseStart stands for the start phase of the funtion
	PhaseStart = "Start"
	// PhaseInProgress stands that function is running
	PhaseInProgress = "InProgress"
	// PhaseEnd stands for the end phase of the function
	PhaseEnd = "End"
)

// Track records all track information
type Track struct {
	Config
	mutex   *sync.Mutex
	callers map[CallerName]time.Time
	tickers map[CallerName]*time.Ticker
}

// New creates a track handler based on custom configuration
func New(config Config) *Track {
	return &Track{
		Config:  config,
		mutex:   &sync.Mutex{},
		callers: make(map[CallerName]time.Time),
		tickers: make(map[CallerName]*time.Ticker),
	}
}

// Default returns a track handler to allow use Start() and End() method
func Default() *Track {
	return &Track{
		Config: Config{
			Debug:           true,
			AsynLog:         true,
			AsynLogInterval: 5,
		},
		mutex:   &sync.Mutex{},
		callers: make(map[CallerName]time.Time),
		tickers: make(map[CallerName]*time.Ticker),
	}
}

// Start records the start of a function
func (t *Track) Start() {
	if t.Debug {
		curCaller := t.callerName()
		t.callers[curCaller] = time.Now()
		if t.AsynLog {
			go t.inProgress(curCaller)
		}
		t.print(PhaseStart, curCaller, 0)

	}
}

func (t *Track) inProgress(s CallerName) {
	if t.Debug {
		t.tickers[s] = time.NewTicker(time.Duration(t.AsynLogInterval) * time.Second)
		for range t.tickers[s].C {
			t.print(PhaseInProgress, s, 0)
		}
	}
}

// End records the end of the function
func (t *Track) End() {
	if t.Debug {
		curCaller := t.callerName()
		if t.AsynLog {
			t.tickers[curCaller].Stop()
			delete(t.tickers, curCaller)
		}
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
	case PhaseInProgress:
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

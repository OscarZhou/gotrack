package track

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	export  bool
	Error   error
}

// New creates a track handler based on custom configuration
func New(config Config) *Track {
	t := &Track{
		Config:  config,
		mutex:   &sync.Mutex{},
		callers: make(map[CallerName]time.Time),
		tickers: make(map[CallerName]*time.Ticker),
	}

	if config.ExportedPath != "" {
		t.export = true
		t.Error = checkPath(config.ExportedPath)
	}
	return t
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

	info := ""
	switch p {
	case PhaseStart:
		info = fmt.Sprintf("%s function:\t%v \n", p, string(s))
	case PhaseInProgress:
		info = fmt.Sprintf("%s function:\t%v \n", p, string(s))
	case PhaseEnd:
		info = fmt.Sprintf("%s function:\t%v took %v \n", p, string(s), elapsed)
	}

	f, err := os.OpenFile(t.ExportedPath, os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		t.Error = err
		return
	}
	f.WriteString(info)
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

func checkPath(p string) error {
	p = filepath.FromSlash(p)
	if strings.ContainsRune(p, os.PathSeparator) {
		lastSlash := strings.LastIndex(p, string(os.PathSeparator))
		if lastSlash > 0 {
			dir := p[0:lastSlash]
			fmt.Println(dir)

			if _, err := os.Stat(p); os.IsNotExist(err) {
				if err := os.MkdirAll(p, os.ModePerm); err != nil {
					return err
				}
			}
		}

		_, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

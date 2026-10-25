// Package debug provides optional logging for the 2048 game.
// Logging is disabled unless the pre-existing log file can be opened.
package debug

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

const Path = "/mnt/sdcard/cubegm/logs/2048.log"

var mu sync.Mutex
var file *os.File

func Init() {
	mu.Lock()
	defer mu.Unlock()
	if file != nil {
		return
	}
	info, err := os.Stat(Path)
	if err != nil || !info.Mode().IsRegular() {
		return
	}
	f, err := os.OpenFile(Path, os.O_WRONLY|os.O_APPEND, 0)
	if err == nil {
		file = f
	}
}

func Logf(format string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	if file == nil {
		return
	}
	_, _ = fmt.Fprintf(file, "[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	_ = file.Sync()
}

func Panic(v any) {
	Logf("PANIC: %v\n%s", v, debug.Stack())
}

func Close() {
	mu.Lock()
	defer mu.Unlock()
	if file != nil {
		_ = file.Close()
		file = nil
	}
}

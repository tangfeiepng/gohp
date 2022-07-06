package log

import (
	"Walker/global"
	"fmt"
	"github.com/gocolly/colly/v2/debug"
	"io"
	"log"
	"os"
	"sync/atomic"
	"time"
)

type ReptileLog struct {
	Output io.Writer
	// Prefix appears at the beginning of each generated log line
	Prefix string
	// Flag defines the logging properties.
	Flag    int
	logger  *log.Logger
	counter int32
	start   time.Time
}

func (l *ReptileLog) Init() error {
	l.counter = 0
	l.start = time.Now()
	if l.Output == nil {
		l.Output = os.Stderr
	}
	l.logger = log.New(l.Output, l.Prefix, l.Flag)
	return nil
}

// Event receives Collector events and prints them to STDERR
func (l *ReptileLog) Event(e *debug.Event) {
	i := atomic.AddInt32(&l.counter, 1)
	str := fmt.Sprintf("[%06d] %d [%6d - %s] %q (%s)\n", i, e.CollectorID, e.RequestID, e.Type, e.Values, time.Since(l.start))
	global.InitLogger.Info(str)
}

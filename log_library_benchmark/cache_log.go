package main

import (
	"io"
	"os"
	"sync"
)

// Messages cache all log messages
// type Messages struct {
// 	Msgs []string
// }

// CacheLogger is logger not flush directly
// Cache first, and flush together later
type CacheLogger struct {
	Out io.Writer

	msgs []string
	mu   sync.Mutex
}

func NewCacheLogger() *CacheLogger {
	msgs := make([]string, 0, 100)
	return &CacheLogger{
		Out:  os.Stderr,
		msgs: msgs,
	}
}

func (l *CacheLogger) AppendMsg(msg string) {
	// msgs := l.bufPool.Get().([]string)
	// defer l.bufPool.Put(msgs)

	l.mu.Lock()
	l.msgs = append(l.msgs, msg)
	l.mu.Unlock()
}

// Flush all cache to out
func (l *CacheLogger) Flush() error {
	// msgs := l.bufPool.Get().([]string)
	// defer l.bufPool.Put(msgs)

	l.mu.Lock()
	defer l.mu.Unlock()
	for _, msg := range l.msgs {
		_, err := l.Out.Write([]byte(msg))
		if err != nil {
			return err
		}
	}

	l.msgs = l.msgs[:0]

	return nil
}

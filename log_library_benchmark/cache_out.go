package main

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// CacheWriter is io.Writer with cache
type CacheWriter struct {
	Out io.Writer

	mu     sync.Mutex
	buf    *bytes.Buffer
	bufHit int32
}

func NewCacheWriter() *CacheWriter {
	return &CacheWriter{
		Out: os.Stderr,
		buf: new(bytes.Buffer),
	}
}

func (cw *CacheWriter) Info(msg string) {

}

func (cw *CacheWriter) AppendMsg(msg string) (int, error) {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	return cw.buf.WriteString(msg)
}

func (cw *CacheWriter) Flush() error {
	cw.mu.Lock()
	_, err := cw.Out.Write(cw.buf.Bytes())
	if err != nil {
		return err
	}

	cw.buf.Reset()
	cw.mu.Unlock()

	return nil
}

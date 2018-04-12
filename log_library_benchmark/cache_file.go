package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type CacheFile struct {
	Out io.Writer

	mu  sync.Mutex
	hit int32
	buf *bytes.Buffer
}

func NewCacheFile() *CacheFile {
	return &CacheFile{
		Out: os.Stderr,
		buf: new(bytes.Buffer),
	}
}

func (cf *CacheFile) Info(msg string) {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	_, err := cf.buf.WriteString(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write buffer error: %v", err)
		return
	}
	cf.hit++

	if cf.hit == 100 {
		_, err := cf.Out.Write(cf.buf.Bytes())
		if err != nil {
			fmt.Fprintf(os.Stderr, "write buffer error: %v", err)
			return
		}

		cf.buf.Reset()
	}
}

func (cf *CacheFile) Flush() {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	if cf.buf.Len() == 0 {
		return
	}

	_, err := cf.Out.Write(cf.buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "write buffer error: %v", err)
		return
	}
}

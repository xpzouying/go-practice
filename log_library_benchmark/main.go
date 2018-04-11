package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

const loopCount = 1000000
const logStr = "this is logging string. this is logging string."

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func logrusOutput() {
	logger := logrus.New()

	f, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()
	logger.Out = f

	begin := time.Now()
	for i := 0; i < loopCount; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Infof(s)
	}
	timeUsed := time.Since(begin)

	fmt.Printf("logrus count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func logrusOutputWithField() {
	logger := logrus.New()

	f, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()
	logger.Out = f

	begin := time.Now()
	for i := 0; i < loopCount; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.WithField("key", "value").Infof(s)
	}
	timeUsed := time.Since(begin)

	fmt.Printf("logrus count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func logrusOutputParallel() {
	logger := logrus.New()
	logger.SetNoLock()

	f, err := os.OpenFile("logrus_parallel.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()
	logger.Out = f

	const subGroupSize = 100
	groupCount := loopCount / subGroupSize
	begin := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < groupCount; i++ {
		wg.Add(1)
		groupNum := i
		go func() {
			for j := 0; j < subGroupSize; j++ {
				cnt := groupNum*subGroupSize + j
				s := fmt.Sprintf("%s, %d", logStr, cnt)
				logger.Infof(s)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	timeUsed := time.Since(begin)

	fmt.Printf("logrus parallel count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

// logrus, output in batch entry
func logrusOutputParallelEntries() {
	logger := logrus.New()
	logger.SetNoLock()

	f, err := os.OpenFile("logrus_parallel_entries.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()
	logger.Out = f

	const subGroupSize = 100
	groupCount := loopCount / subGroupSize
	begin := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < groupCount; i++ {
		wg.Add(1)
		groupNum := i
		go func() {
			// buf := bufPool.Get().(*bytes.Buffer)
			buf := new(bytes.Buffer)
			buf.Reset()

			for j := 0; j < subGroupSize; j++ {
				cnt := groupNum*subGroupSize + j

				entry := logrus.NewEntry(logger)
				entry.Time = time.Now()
				entry.Message = fmt.Sprintf("%s, %d", logStr, cnt)
				entry.Level = logrus.InfoLevel

				fmtMsg, err := entry.Logger.Formatter.Format(entry)
				if err != nil {
					panic(err)
				}

				_, err = buf.WriteString(string(fmtMsg))
				if err != nil {
					panic(err)
				}
			}

			_, err := logger.Out.Write(buf.Bytes())
			if err != nil {
				panic(err)
			}

			wg.Done()
		}()
	}
	wg.Wait()
	timeUsed := time.Since(begin)

	fmt.Printf("logrus parallel count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func logrusOutputParallelInBuffer() {
	logger := logrus.New()
	logger.SetNoLock()

	f, err := os.OpenFile("logrus_parallel_in_buffer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()
	logger.Out = f

	// bufPool for caching logging
	// bufPool := sync.Pool{
	// 	New: func() interface{} {
	// 		return new(bytes.Buffer)
	// 	},
	// }

	const subGroupSize = 100
	groupCount := loopCount / subGroupSize
	begin := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < groupCount; i++ {
		wg.Add(1)
		groupNum := i
		go func() {
			// buf := bufPool.Get().(*bytes.Buffer)
			buf := new(bytes.Buffer)
			buf.Reset()

			for j := 0; j < subGroupSize; j++ {
				cnt := groupNum*subGroupSize + j
				buf.WriteString(fmt.Sprintf("%s, %d", logStr, cnt))
				buf.WriteString("\n")
			}
			logger.Infof(buf.String())
			wg.Done()
		}()
	}
	wg.Wait()
	timeUsed := time.Since(begin)

	fmt.Printf("logrus parallel count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func zerologOutput() {
	f, err := os.OpenFile("zero-log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()

	logger := zerolog.New(f)

	begin := time.Now()
	for i := 0; i < loopCount; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info().Msg(s)
	}
	timeUsed := time.Since(begin)

	fmt.Printf("zerolog count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func uberZapOutput() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"zap_test.log"}

	logger, err := cfg.Build()
	checkError(err)

	begin := time.Now()
	for i := 0; i < loopCount; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info(s)
	}
	logger.Sync()

	timeUsed := time.Since(begin)

	fmt.Printf("uber zap count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func stdLogOutput() {
	f, err := os.OpenFile("stdlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	checkError(err)
	defer f.Close()

	log.SetOutput(f)

	begin := time.Now()
	for i := 0; i < loopCount; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		log.Println(s)
	}
	timeUsed := time.Since(begin)

	fmt.Printf("stdlog count:%d, time_used: %v, time_per_op: %v\n", loopCount, timeUsed, timeUsed/loopCount)
}

func main() {
	// logrusOutput()
	// logrusOutputParallel()
	logrusOutputParallelEntries()
	// logrusOutputParallelInBuffer()

	// zerologOutput()

	// uberZapOutput()

	// stdLogOutput()
}

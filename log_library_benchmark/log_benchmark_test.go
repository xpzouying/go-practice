package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func BenchmarkLogrusDefaultConfig(b *testing.B) {
	logger := logrus.New()

	f, err := ioutil.TempFile("", "logrus_log")
	assert.NoError(b, err)
	defer f.Close()
	logger.Out = f

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Infof(s)
	}
}

func BenchmarkLogrusDefaultConfigWithBufio(b *testing.B) {
	logger := logrus.New()

	f, err := ioutil.TempFile("", "logrus_log")
	assert.NoError(b, err)
	defer f.Close()

	cacheOut := bufio.NewWriter(f)
	logger.Out = cacheOut

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Infof(s)
	}
	assert.NoError(b, cacheOut.Flush())
}

func BenchmarkLogrusDisable(b *testing.B) {
	f, err := ioutil.TempFile("", "logrus_log")
	assert.NoError(b, err)
	defer f.Close()

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}
	logger.Out = f

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Infof(s)
	}
}

func BenchmarkLogrusWithField(b *testing.B) {
	logger := logrus.New()

	f, err := ioutil.TempFile("", "logrus_log")
	assert.NoError(b, err)
	defer f.Close()
	logger.Out = f

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.WithField("key", "value").Infof(s)
	}
}

func BenchmarkZerolog(b *testing.B) {
	f, err := ioutil.TempFile("", "zero_log")
	assert.NoError(b, err)
	defer f.Close()

	logger := zerolog.New(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info().Msg(s)
	}
}

func BenchmarkZerologWithBufio(b *testing.B) {
	f, err := ioutil.TempFile("", "zero_log")
	assert.NoError(b, err)
	defer f.Close()

	cacheOut := bufio.NewWriterSize(f, 102400)

	logger := zerolog.New(cacheOut)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info().Msg(s)
	}
	assert.NoError(b, cacheOut.Flush())
}

func BenchmarkUberZap(b *testing.B) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"zap_test.log"}

	logger, err := cfg.Build()
	assert.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info(s)
	}
	logger.Sync()
}

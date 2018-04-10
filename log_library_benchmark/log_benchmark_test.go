package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func BenchmarkLogrus(b *testing.B) {
	logger := logrus.New()

	f, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	assert.NoError(b, err)
	defer f.Close()
	logger.Out = f

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Infof(s)
	}
}

func BenchmarkZerolog(b *testing.B) {
	f, err := os.OpenFile("zero-log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	assert.NoError(b, err)
	defer f.Close()

	logger := zerolog.New(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s, %d", logStr, i)
		logger.Info().Msg(s)
	}
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

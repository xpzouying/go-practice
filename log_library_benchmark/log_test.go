package main

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestUberZap(t *testing.T) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"zap_test.log"}

	logger, err := cfg.Build()
	assert.NoError(t, err)

	logger.Info(logStr)
}

func TestZerolog(t *testing.T) {
	f, err := os.OpenFile("zero_test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	assert.NoError(t, err)
	defer f.Close()

	logger := zerolog.New(f)

	logger.Info().Msg(logStr)
}

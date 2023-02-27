package main

import (
	"runtime"

	"github.com/K0STYAa/vk_iproto/internal/app"
	"github.com/K0STYAa/vk_iproto/pkg/logger"
)

const (
	GoMaxProcsLim = 4
)

func main() {
	runtime.GOMAXPROCS(GoMaxProcsLim)
	logger.LogStart()

	app.Run()
}

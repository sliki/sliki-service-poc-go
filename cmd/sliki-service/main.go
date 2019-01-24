package main

import (
	"fmt"
	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

func main() {
	log, _ = zap.NewDevelopment()
	log.Info("log foo")
	fmt.Println("foo")
}

package main

import (
	"os"

	"github.com/westcottian/buyrent-breakevencalculator/buyrentcalculator"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

func main() {
	// Create separate main instead of doing the actual code here
	// since os.Exit can not handle `defer`. DON'T call `os.Exit` in the any other place.
	os.Exit(realMain(os.Args))
}

// realMain is start point of service. We initialize required client
// (DB or external services) and logger and pass it to the server struct.
// nolint: gocyclo, funlen
func realMain(_ []string) int {

	// configure our calculator service
	calculatorService := buyrentcalculator.NewService()

	// Create new gRPC server. gRPC server handles core function this service.
	grpcServer, err := grpc.New(calculatorService)
	})
	if err != nil {
		logger.Error("failed to setup new gRPC server", zap.Error(err))
		return exitError
	}

}

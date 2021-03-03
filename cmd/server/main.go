package main

import (
	"os"

	"github.com/westcottian/buyrent-breakevencalculator/buyrentcalculator"
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

}

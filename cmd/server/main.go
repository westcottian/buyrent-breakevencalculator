package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/westcottian/buyrent-breakevencalculator/buyrentcalculator"
	"github.com/westcottian/buyrent-breakevencalculator/di"
	"github.com/westcottian/buyrent-breakevencalculator/grpc"
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
	c := di.NewContainer()
	env := c.InjectConfig()
	// configure our calculator service
	calculatorService := buyrentcalculator.NewService()

	// Create new gRPC server. gRPC server handles core function this service.
	grpcServer, err := grpc.New(calculatorService)
	if err != nil {
		fmt.Println("failed to setup new gRPC server")
		return exitError
	}
	// Start listening port for gRPC server.
	grpcAddr := fmt.Sprintf(":%d", env.GRPCPort)
	fmt.Println("grpc server listening")
	grpcLn, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		fmt.Println("failed to listen gRPC port")
		return exitError
	}

	// Start serving gRPC and HTTP server in different goroutines.
	// And also start to fetch public key from authority service.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return grpcServer.Serve(grpcLn) })

	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		fmt.Println("received SIGTERM, exiting server gracefully")
	case <-ctx.Done():
	}

	// Gracefully shut down both gRPC and HTTP server in parallel.
	// before initiating the graceful shutdown
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		fmt.Println("gracefully shutdown gRPC server")
		grpcServer.GracefulStop()
		fmt.Println("completed to gracefully shutdown gRPC server")
	}()

	<-doneCh

	cancel()
	if err := wg.Wait(); err != nil {
		fmt.Println("unhandled error received")
		return exitError
	}

	return exitOK

}

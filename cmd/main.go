package main

import (
    "net"
    "os"
    "os/signal"
    "syscall"
    "github.com/aldoclg/grpc/app"
    greeting_v1 "github.com/aldoclg/grpc/package/proto/v1/service"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    _ "google.golang.org/grpc/health"
    _ "google.golang.org/grpc/health/grpc_health_v1"
)

var (
    listener net.Listener
    server *grpc.Server
    logger *zap.Logger
)

func main() {
    logger, _  := zap.NewProduction()
    defer logger.Sync()

    initListener()

    server = grpc.NewServer()

    greeting_v1.RegisterGreeterServiceServer(server, &app.GreeterServiceServer{})

    logger.Info("Handlers registered")

    go signalsListener(server)

    logger.Info("Starting gRPC server")

    if err := server.Serve(listener); err != nil {
        logger.Panic("Failed to start gRPC  server", zap.Error(err))
    }

}
func initListener() {
    var err error
    addr := "localhost:50051"

    listener, err = net.Listen("tcp", addr)

    if err != nil {
        logger.Panic("Failed to listen", zap.Error(err))
    }

    // logger.Info("Started listening...")
    return
}

func signalsListener(server *grpc.Server) {
    sigs := make(chan os.Signal, 1)

    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
    _ = <-sigs

    // logger.Info("Gracefully stopping server...")
    server.GracefulStop()
}

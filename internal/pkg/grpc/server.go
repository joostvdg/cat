package grpc

import (
    "context"
    "net"
    "os"
    "os/signal"

    "google.golang.org/grpc"

    "github.com/joostvdg/cat/pkg/api/v1"

    joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"
)

func init() {
    // Log as JSON instead of the default ASCII formatter.
    // log.SetFormatter(&log.JSONFormatter{})

    log.SetReportCaller(true)
    log.SetLevel(log.InfoLevel)
    log.SetFormatter(&joonix.FluentdFormatter{})
}

// RunServer runs gRPC service to publish Application service
func RunServer(ctx context.Context, v1API v1.ApplicationServiceServer, port string) error {
    listen, err := net.Listen("tcp", ":"+port)
    if err != nil {
        return err
    }

    // register service
    server := grpc.NewServer()
    v1.RegisterApplicationServiceServer(server, v1API)

    // graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for range c {
            // sig is a ^C, handle it
            log.Println("shutting down gRPC server...")

            server.GracefulStop()

            <-ctx.Done()
        }
    }()

    // start gRPC server
    log.WithFields(log.Fields{
        "port": port,
    }).Info("Starting grpc server")
    return server.Serve(listen)
}

package cmd

import (
	"context"
    "fmt"

    joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"

    "github.com/joostvdg/cat/internal/pkg/persistence"
	"github.com/joostvdg/cat/internal/pkg/grpc"
	"github.com/joostvdg/cat/pkg/service/v1"
	v1_api "github.com/joostvdg/cat/pkg/api/v1"
)

func init() {
    // Log as JSON instead of the default ASCII formatter.
    // log.SetFormatter(&log.JSONFormatter{})

    log.SetReportCaller(true)
    log.SetLevel(log.InfoLevel)
    log.SetFormatter(&joonix.FluentdFormatter{})
}

// StartGRPCServer runs gRPC server and HTTP gateway
func StartGRPCServer(persistenceBackendType string, port string) error {
	ctx := context.Background()

    persistenceBackend, err := persistence.InitPersistenceBackend(persistenceBackendType)
    if err != nil {
        return fmt.Errorf("%s is not a supported persistence backend", persistenceBackendType)
    }
    persistenceBackend.Add(v1_api.Application{
        Name:        "Jenkins",
        Description: "Jenkins, the most awesome CI engine",
        Namespace:   "CI",
        ArtifactIDs: []string{"https://registry.hub.docker.com/library/jenkins@sha256:81040e35ee59322a02f67ca2584f814d543d5f2f5d361fb8bf4f9e0046f3e809"},
        Sources:     []string{"https://github.com/jenkinsci/jenkins.git"},
        Labels:      []*v1_api.Label{{Key: "Category", Value: "BuildTool"}},
        Annotations: []*v1_api.Annotation{{Key: "MetricsGroup", Value: "CI", Origin: "com.github.joostvdg"}},
    })
    log.WithFields(log.Fields{
        "type": persistenceBackend.PersistenceType(),
        "items": len(persistenceBackend.GetAll()),
    }).Info("Initialized persistence")

    v1API := v1.NewApplicationServiceServer(persistenceBackend)

	return grpc.RunServer(ctx, v1API, port)
}

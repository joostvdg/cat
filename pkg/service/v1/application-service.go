package v1

import (
    "context"
    "fmt"
    "github.com/joostvdg/cat/internal/pkg/persistence"
    "github.com/joostvdg/cat/pkg/api/v1"
    joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func init() {
    // Log as JSON instead of the default ASCII formatter.
    // log.SetFormatter(&log.JSONFormatter{})

    log.SetReportCaller(true)
    log.SetLevel(log.InfoLevel)
    log.SetFormatter(&joonix.FluentdFormatter{})
}

const (
    // apiVersion is version of API is provided by server
    apiVersion = "v1"
)

func NewApplicationServiceServer(persistenceBackend persistence.PersistenceBackend) v1.ApplicationServiceServer {
    return &applicationServiceServer{
        PersistenceBackend: persistenceBackend,
    }
}

type applicationServiceServer struct {
    PersistenceBackend persistence.PersistenceBackend
}

// checkAPI checks if the API version requested by client is supported by server
func (as *applicationServiceServer) checkAPI(api string) error {
    // API version is "" means use current version of the service
    if len(api) > 0 {
        if apiVersion != api {
            return status.Errorf(codes.Unimplemented,
                "unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
        }
    }
    return nil
}

func (as *applicationServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
    // check if the API version requested by client is supported by server
    if err := as.checkAPI(req.Api); err != nil {
        return nil, err
    }

    app := *req.Application
    if app.Uuid != "" && as.PersistenceBackend.Exists(app.Uuid) {
        log.Warnf("Application with uuid=%v already exists", app.Uuid)
        return &v1.CreateResponse {
            Api: apiVersion,
            Uuid:  "",
        }, fmt.Errorf("cannot create, application already exists")
    } else {
        if app.Uuid == "" {
            as.PersistenceBackend.Add(app)
            log.Infof("Created new application entry, uuid=%v, name=%v, current count=%v", app.Uuid, app.Name, len(as.PersistenceBackend.GetAll()))
        }
    }

    return &v1.CreateResponse {
        Api: apiVersion,
        Uuid:  app.Uuid,
    }, nil
}

func (as *applicationServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error){
    if err := as.checkAPI(req.Api); err != nil {
        return nil, err
    }

    app := as.PersistenceBackend.GetOne(req.Uuid)
    return &v1.ReadResponse{
        Api: apiVersion,
        Application: &app,
    }, nil
}

func (as *applicationServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error){
    if err := as.checkAPI(req.Api); err != nil {
        return nil, err
    }

    // TODO: not yet supported
    return &v1.UpdateResponse{}, nil
}

func (as *applicationServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
    if err := as.checkAPI(req.Api); err != nil {
        return nil, err
    }

    as.PersistenceBackend.Remove(req.Uuid)
    return &v1.DeleteResponse{}, nil
}

func (as *applicationServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
    if err := as.checkAPI(req.Api); err != nil {
        return nil, err
    }

    apps := as.PersistenceBackend.GetAll()
    applications :=  make([]*v1.Application, len(apps))
    for i,app := range apps {
        applications[i] = &app
    }
    log.Info("Retrieving all applications, found: ", len(applications))

    return &v1.ReadAllResponse{
        Api: apiVersion,
        Applications: applications,
    }, nil
}

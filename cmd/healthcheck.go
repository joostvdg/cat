package cmd

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"

    joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"
)

func init() {
    log.SetReportCaller(true)
    log.SetLevel(log.InfoLevel)
    log.SetFormatter(&joonix.FluentdFormatter{})
}

func StartHealthCheck(endpointType string) {
    // ignore for now
    StartServer("80")
}


type Server struct {
    mux    *http.ServeMux
}

// StartServer Starts the web server on the given port with the given data
func StartServer(port string) {
    router := mux.NewRouter()
    router.HandleFunc("/healthz", HandleHealthCheck)
    listenAddress := fmt.Sprintf(":%s", port)
    server := &http.Server{Addr: listenAddress, Handler: router}
    log.WithFields(log.Fields{
        "port": port,
        "endpoint": "/healthz",
    }).Info("Initializing http health check")
    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}

type HealthStatus struct {
    Status      string
    Description string
}

// HandleHealthCheck is the handler function for serving the health checks
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
    status := HealthStatus{Status: "OK", Description: "All is good."}
    json.NewEncoder(w).Encode(status)
}

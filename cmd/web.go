package cmd

import (
	"fmt"
    "net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/google/uuid"

	joonix "github.com/joonix/log"
    log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

    "github.com/joostvdg/cat/internal/pkg/persistence"
    "github.com/joostvdg/cat/internal/pkg/web"
    "github.com/joostvdg/cat/pkg/api/v1"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&joonix.FluentdFormatter{})
}

var persistenceBackend persistence.PersistenceBackend

func StartWebserver(port string, persistenceBackendType string) {
	log.WithFields(log.Fields{
		"type": "in-memory",
	}).Info("Initializing persistence")

	err := initPersistence(persistenceBackendType)
	if err != nil {
		log.Error("Could not initialize persistence backend", "Error", err)
	}

	log.Info("Initializing web server")
	initWebServer(port)
}

func initPersistence(backendType string) error {
	persistenceTemp, err := persistence.InitPersistenceBackend(backendType)
	if err != nil {
		return fmt.Errorf("%s is not a supported persistence backend", backendType)
	}
	persistenceBackend = persistenceTemp

	persistenceBackend.Add(v1.Application{
		Name:        "Maven Demo Library",
		Description: "A small Maven Java library for demo purposes",
		Uuid:        uuid.New().String(),
		Namespace:   "joostvdg",
		ArtifactIDs: []string{"gav://com.github.joostvdg.demo:maven-demo-lib:0.1.1"},
		Sources:     []string{"https://github.com/joostvdg/maven-demo-lib.git"},
		Labels:      []*v1.Label{{Key: "Category", Value: "BuildTool"}},
		Annotations: []*v1.Annotation{{Key: "MetricsGroup", Value: "CI", Origin: "com.github.joostvdg"}},
	})

	persistenceBackend.Add(v1.Application{
		Name:        "Jenkins",
		Description: "Jenkins, the most awesome CI engine",
		Uuid:        uuid.New().String(),
		Namespace:   "CI",
		ArtifactIDs: []string{"https://registry.hub.docker.com/library/jenkins@sha256:81040e35ee59322a02f67ca2584f814d543d5f2f5d361fb8bf4f9e0046f3e809"},
		Sources:     []string{"https://github.com/jenkinsci/jenkins.git"},
		Labels:      []*v1.Label{{Key: "Category", Value: "BuildTool"}},
		Annotations: []*v1.Annotation{{Key: "MetricsGroup", Value: "CI", Origin: "com.github.joostvdg"}},
	})
	return nil
}

func initWebServer(port string) {
	// Setup
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Server.Addr = ":" + port
	e.GET("/users/:id", getUser)
	e.GET("/applications", web.GetApplications)
	e.PUT("/applications", web.PostApplication)

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &web.CustomContext{
				c,
				persistenceBackend,
			}
			return h(cc)
		}
	})

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Starting server")
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

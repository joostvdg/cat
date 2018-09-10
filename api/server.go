package api

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/google/uuid"
	"github.com/joostvdg/cat/application"
	"github.com/joostvdg/cat/persistence"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"net/http"
)

func initPersistence() {
	persistence.InitMemoryMap()
	persistence.Add(application.Application{
		Name:        "Maven Demo Library",
		Description: "A small Maven Java library for demo purposes",
		UUID:        uuid.New().String(),
		Namespace:   "joostvdg",
		ArtifactIDs: []string{"gav://com.github.joostvdg.demo:maven-demo-lib:0.1.1"},
		Sources:     []string{"https://github.com/joostvdg/maven-demo-lib.git"},
	})

	persistence.Add(application.Application{
		Name:        "Jenkins",
		Description: "Jenkins, the most awesome CI engine",
		UUID:        uuid.New().String(),
		Namespace:   "CI",
		ArtifactIDs: []string{"https://registry.hub.docker.com/library/jenkins@sha256:81040e35ee59322a02f67ca2584f814d543d5f2f5d361fb8bf4f9e0046f3e809"},
		Sources:     []string{"https://github.com/jenkinsci/jenkins.git"},
	})
}

func Serve(port string) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	sugar.Infow("Initializing persistence",
		"type", "in-memory",
	)
	initPersistence()

	// Setup
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Server.Addr = ":" + port
	e.GET("/users/:id", getUser)
	e.GET("/applications", GetApplications)
	e.POST("/applications", PostApplication)

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Serve it like a boss
	sugar.Infow("Starting server",
		"port", port,
	)
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

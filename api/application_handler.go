package api

import (
    "github.com/google/uuid"
    "github.com/joostvdg/cat/application"
    "github.com/joostvdg/cat/persistence"
    "github.com/labstack/echo"
    "net/http"
)

func GetApplications(c echo.Context) error {
	return c.JSON(http.StatusOK, persistence.GetAll())
}

func PostApplication(c echo.Context) (err error) {
	app := new(application.Application)
	if err = c.Bind(app); err != nil {
		return
	}

	if persistence.Exists(*app) {
		return c.JSON(http.StatusSeeOther, persistence.GetOne(app.UUID))
	}

	returnStatus := http.StatusAccepted
    if app.UUID == "" {
        app.UUID = uuid.New().String()
        returnStatus = http.StatusCreated
    }
    persistence.Add(*app)
	return c.JSON(returnStatus, app)
}

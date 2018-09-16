package api

import (
	"github.com/google/uuid"
	"github.com/joostvdg/cat/application"
	"github.com/labstack/echo"
	"net/http"
)

func GetApplications(c echo.Context) error {
	cc := c.(*CustomContext)
	return c.JSON(http.StatusOK, cc.PersistenceBackend.GetAll())
}

func PostApplication(c echo.Context) (err error) {
	cc := c.(*CustomContext)
	app := new(application.Application)
	if err = c.Bind(app); err != nil {
		return
	}

	if cc.PersistenceBackend.Exists(*app) {
		return c.JSON(http.StatusSeeOther, cc.PersistenceBackend.GetOne(app.UUID))
	}

	returnStatus := http.StatusAccepted
	if app.UUID == "" {
		app.UUID = uuid.New().String()
		returnStatus = http.StatusCreated
	}
	cc.PersistenceBackend.Add(*app)
	return c.JSON(returnStatus, app)
}

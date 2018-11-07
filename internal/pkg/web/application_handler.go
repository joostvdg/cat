package web

import (
	"github.com/joostvdg/cat/pkg/api/v1"
	"github.com/labstack/echo"
	"net/http"
)

func GetApplications(c echo.Context) error {
	cc := c.(*CustomContext)
	return c.JSON(http.StatusOK, cc.PersistenceBackend.GetAll())
}

func PostApplication(c echo.Context) (err error) {
	cc := c.(*CustomContext)
	app := new(v1.Application)
	if err = c.Bind(app); err != nil {
		return
	}

	if cc.PersistenceBackend.Exists(app.Uuid) {
		return c.JSON(http.StatusSeeOther, cc.PersistenceBackend.GetOne(app.Uuid))
	}

	returnStatus := http.StatusCreated
	_, addErr := cc.PersistenceBackend.Add(*app)
	if addErr != nil {
        returnStatus = http.StatusBadRequest
    }
	return c.JSON(returnStatus, app)
}

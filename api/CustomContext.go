package api

import (
	"github.com/joostvdg/cat/persistence"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	PersistenceBackend persistence.PersistenceBackend
}

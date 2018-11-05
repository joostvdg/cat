package web

import (
	"github.com/joostvdg/cat/internal/pkg/persistence"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	PersistenceBackend persistence.PersistenceBackend
}

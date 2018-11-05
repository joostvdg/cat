package persistence

import (
	"github.com/joostvdg/cat/pkg/application"
)

type memory struct {
	Applications map[string]application.Application
}

func (m *memory) GetAllIds() []string {
	keys := make([]string, 0, len(m.Applications))
	for key := range m.Applications {
		keys = append(keys, key)
	}
	return keys
}

func (m *memory) GetAll() []application.Application {
	apps := make([]application.Application, 0, len(m.Applications))
	for _, app := range m.Applications {
		apps = append(apps, app)
	}
	return apps
}

func (m *memory) Add(app application.Application) {
	m.Applications[app.UUID] = app
}

func (m *memory) Exists(app application.Application) bool {
	_, ok := m.Applications[app.UUID]
	return ok
}

func (m *memory) GetOne(uuid string) application.Application {
	return m.Applications[uuid]
}

func (m *memory) Remove(app application.Application) {
	delete(m.Applications, app.UUID)
}

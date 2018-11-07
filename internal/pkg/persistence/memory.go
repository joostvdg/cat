package persistence

import (
    "fmt"
    "github.com/google/uuid"
    "github.com/joostvdg/cat/pkg/api/v1"
)

type memory struct {
	Applications map[string]v1.Application
}


func (m *memory) PersistenceType() string {
    return "in-memory"
}

func (m *memory) GetAllIds() []string {
	keys := make([]string, 0, len(m.Applications))
	for key := range m.Applications {
		keys = append(keys, key)
	}
	return keys
}

func (m *memory) GetAll() []v1.Application {
	apps := make([]v1.Application, 0, len(m.Applications))
	for _, app := range m.Applications {
		apps = append(apps, app)
	}
	return apps
}

func (m *memory) Add(app v1.Application) (string, error) {
    if app.Uuid != "" {
        return "", fmt.Errorf("cannot add application, already has a UUID")
    }
    app.Uuid = uuid.New().String()
    m.Applications[app.Uuid] = app
    return app.Uuid, nil
}

func (m *memory) Exists(uuid string) bool {
	_, ok := m.Applications[uuid]
	return ok
}

func (m *memory) GetOne(uuid string) v1.Application {
	return m.Applications[uuid]
}

func (m *memory) Remove(uuid string) bool {
    removed := false
    if m.Exists(uuid) {
        delete(m.Applications, uuid)
        removed = m.Exists(uuid)
    }
	return removed
}

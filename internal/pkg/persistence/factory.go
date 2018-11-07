package persistence

import (
	"fmt"
    "github.com/joostvdg/cat/pkg/api/v1"
)

func initMemoryMap() PersistenceBackend {
	m := memory{
		Applications: make(map[string]v1.Application),
	}
	return &m
}

func InitPersistenceBackend(persistenceBackend string) (PersistenceBackend, error) {
	switch persistenceBackend {
	case MEM:
		return initMemoryMap(), nil
	default:
		return &Empty{}, fmt.Errorf("%s is not a supported persistence backend", persistenceBackend)
	}
}

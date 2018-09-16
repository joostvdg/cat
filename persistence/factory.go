package persistence

import (
	"errors"
	"fmt"
	"github.com/joostvdg/cat/application"
)

func initMemoryMap() PersistenceBackend {
	m := memory{
		Applications: make(map[string]application.Application),
	}
	return &m
}

func InitPersistenceBackend(persistenceBackend string) (PersistenceBackend, error) {
	switch persistenceBackend {
	case MEM:
		return initMemoryMap(), nil
	default:
		return &Empty{}, errors.New(fmt.Sprint("%s is not a supported persistence backend", persistenceBackend))
	}
}

package persistence

import (
    "github.com/joostvdg/cat/pkg/api/v1"
)

const (
	MEM  string = "mem"
	ETCD string = "etcd"
)

type PersistenceBackend interface {
	// functions from memory
	// memory should be the first implementation
	// ETCD / Redis / Memory

	PersistenceType() string
	GetAllIds() []string
	GetAll() []v1.Application
	Add(app v1.Application) (string, error)
    Exists(uuid string) bool
	GetOne(uuid string) v1.Application
    Remove(uuid string) bool
}

type Empty struct{}

func (e *Empty) PersistenceType() string {
    return "Empty"
}

func (e *Empty) GetAllIds() []string {
	return make([]string, 0)
}

func (e Empty) GetAll() []v1.Application {
	return make([]v1.Application, 0)
}

func (e Empty) Add(app v1.Application) (string, error) { return "", nil}

func (e Empty) Exists(uuid string) bool { return false }

func (e Empty) GetOne(uuid string) v1.Application { return v1.Application{} }

func (e Empty) Remove(uuid string) bool { return false }

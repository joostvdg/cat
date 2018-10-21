package persistence

import "github.com/joostvdg/cat/application"

const (
	MEM  string = "mem"
	ETCD string = "etcd"
)

type PersistenceBackend interface {
	// functions from memory
	// memory should be the first implementation
	// ETCD / Redis / Memory

	GetAllIds() []string
	GetAll() []application.Application
	Add(app application.Application)
	Exists(app application.Application) bool
	GetOne(uuid string) application.Application
	Remove(app application.Application)
}

type Empty struct{}

func (e *Empty) GetAllIds() []string {
	return make([]string, 0)
}

func (e Empty) GetAll() []application.Application {
	return make([]application.Application, 0)
}

func (e Empty) Add(app application.Application) {}

func (e Empty) Exists(app application.Application) bool { return false }

func (e Empty) GetOne(uuid string) application.Application { return application.Application{} }

func (e Empty) Remove(app application.Application) {}

package persistence

import (
    "github.com/joostvdg/cat/application"
)

var applications map[string]application.Application

func InitMemoryMap() {
	applications = make(map[string]application.Application)
}

// TODO: if query is filter on keys, we only give the ids?!
//func GetAllIds() []string {
//
//}

func GetAll() []application.Application {
	apps := make([]application.Application, 0, len(applications))
	for _,app := range applications {
		apps = append(apps, app)
	}
	return apps
}

func Add(app application.Application) {
	applications[app.UUID] = app
}

func Exists(app application.Application) bool {
	_, ok := applications[app.UUID]
	return ok
}

func GetOne(uuid string) application.Application {
	return applications[uuid]
}

func Remove(app application.Application) {
	delete(applications, app.UUID)
}

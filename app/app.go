package app

import (
	"github.com/VanOODev/education_web_server/adapters/controllers/http/stdrouter"
	"github.com/VanOODev/education_web_server/usecase/notesd"
)

type App struct {
	uc *notesd.NotesD
	sv *stdrouter.Server
}

func NewApp() (a *App, err error) {
	// TODO: обработать конфигурацию, создать все необходимые объекты. Запустить приложение

	a = &App{}
	return a, nil
}

func (a *App) Close() {
	a.uc.Close()
	a.sv.Close()
}

package main

import "sync"

type application struct {
	CfgProvider ConfigProvider
}

var App *application
var appOnce sync.Once

type AppOption func(app *application) error

func InitApplication(opts ...AppOption) error {
	var err error
	appOnce.Do(func() {
		App = &application{
			CfgProvider: NewInMemoryConfigProvider(),
		}
		for _, opt := range opts {
			err = opt(App)
			if err != nil {
				return
			}
		}
	})
	return err
}

func WithCfgProvider(cfg ConfigProvider) AppOption {
	return func(app *application) error {
		app.CfgProvider = cfg
		return nil
	}
}

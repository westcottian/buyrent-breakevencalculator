package di

import (
	"github.com/westcottian/buyrent-breakevencalculator/config"
)

type (
	Container struct {
		cache struct {
			// config
			config *config.Env
		}
	}
)

func NewContainer() *Container {
	container := &Container{}

	// Read configurations from environmental variables.
	container.InjectConfig()

	return container
}

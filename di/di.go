package di

import (
	"fmt"
	"github.com/westcottian/buyrent-breakevencalculator/config"
	"os"
)

func (c *Container) InjectConfig() *config.Env {
	if c.cache.config != nil {
		return c.cache.config
	}

	var err error
	c.cache.config, err = config.ReadFromEnv()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup config: %s\n", err)
		panic(err)
	}

	return c.cache.config
}

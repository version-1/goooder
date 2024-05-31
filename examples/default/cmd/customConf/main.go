package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
)

type config struct {
	seeders []gdconfig.Seeder
}

func (c config) Connstr() string {
	return "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable"
}

func (c config) TemplatePath() string {
	return ""
}

func (c config) Seeders() []gdconfig.Seeder {
	return c.seeders
}

func main() {
	conf := config{
		seeders: development.NewSeed().Seeders(),
	}

	executor := goooder.NewSeedExecutor(conf)
	executor.Run()
}

package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
)

type config struct{}

func (c config) Connstr() string {
	return "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable"
}

func (c config) TemplatePath() string {
	return ""
}

func (c config) Seeders() []gdconfig.Seeder {
	s := development.NewSeed()
	return s.Seeders()
}
func main() {
	conf := config{}

	executor := goooder.NewSeedExecutor(conf)
	executor.Run("Seed_0000010_CreateUsers")
}

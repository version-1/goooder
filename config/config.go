package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// check if EnvConfig implements Config interface
var _ Config = (*EnvConfig)(nil)

type Seeder interface {
	Exec(tx *sqlx.Tx) error
}

type Renderer interface {
	RenderTemplate(filename string) (string, error)
}

type Config interface {
	Connstr() string
	Seeders() []Seeder
}

type EnvConfig struct {
	connstr string
	seeders []Seeder
}

func (e EnvConfig) Connstr() string {
	return e.connstr
}

func (e EnvConfig) Seeders() []Seeder {
	return e.seeders
}

func FromEnv() *EnvConfig {
	mustLoadEnv()

	return &EnvConfig{
		connstr: os.Getenv("DATABASE_CONNSTR"),
	}
}

func (e *EnvConfig) SetSeeders(seeders []Seeder) {
	e.seeders = seeders
}

func mustLoadEnv() {
	var err error
	var envfile string
	env := os.Getenv("GOOODER_ENV")
	if env == "" {
		err = godotenv.Load()
		envfile = ".env"
	} else {
		envfile = ".env." + env
		err = godotenv.Load(envfile)
	}

	if err != nil {
		log.Fatalf("Error loading %s\n", envfile)
	}
}

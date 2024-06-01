package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	connstr string
	seeders []Seeder
	logger  Logger
}

func (e EnvConfig) Connstr() string {
	return e.connstr
}

func (e EnvConfig) Seeders() []Seeder {
	return e.seeders
}

func (e *EnvConfig) Logger() Logger {
	if e.logger == nil {
		l := &DefaultLogger{}
		l.SetLogLevel(strings.ToLower(os.Getenv("LOG_LEVEL")))

		return l
	}
	return e.logger
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

func (e *EnvConfig) SetLogger(logger Logger) {
	e.logger = logger
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

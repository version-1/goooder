package config

import (
	"github.com/jmoiron/sqlx"
)

// check if EnvConfig implements Config interface
var _ Config = (*EnvConfig)(nil)

type Seeder interface {
	Exec(tx *sqlx.Tx) error
}

type Renderer interface {
	RenderTemplate(filename string) (string, error)
}

type Logger interface {
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type Config interface {
	Connstr() string
	Seeders() []Seeder
	Logger() Logger
}

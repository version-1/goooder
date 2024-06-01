package goooder

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	gdconfig "github.com/version-1/goooder/config"
)

type config struct{}

func (c config) Connstr() string {
	return "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable"
}

func (c config) Logger() gdconfig.Logger {
	return gdconfig.DefaultLogger{}
}

func (c config) Seeders() []gdconfig.Seeder {
	return []gdconfig.Seeder{
		&seed_001{},
		&seed_002{},
		&seed_003{},
	}
}

var called = []string{}

type seed_001 struct{}
type seed_002 struct{}
type seed_003 struct{}

func (s *seed_001) Exec(_ *sqlx.Tx) error {
	called = append(called, "seed_001")
	return nil
}

func (s *seed_002) Exec(_ *sqlx.Tx) error {
	called = append(called, "seed_002")
	return nil
}

func (s *seed_003) Exec(_ *sqlx.Tx) error {
	called = append(called, "seed_003")
	return nil
}

func TestSingleSeed(t *testing.T) {
	conf := config{}
	executor := NewSeedExecutor(conf)
	executor.Run("seed_001")

	assert.Equal(t, []string{"seed_001"}, called)
}

func TestAllSeed(t *testing.T) {
	called = []string{}
	conf := config{}
	executor := NewSeedExecutor(conf)
	executor.Run()

	assert.Equal(t, []string{"seed_001", "seed_002", "seed_003"}, called)
}

func TestEnvConfigSeed(t *testing.T) {
	called = []string{}
	conf := gdconfig.FromEnv()
	seeders := []gdconfig.Seeder{
		&seed_001{},
		&seed_002{},
		&seed_003{},
	}

	conf.SetSeeders(seeders)

	executor := NewSeedExecutor(conf)
	executor.Run()

	assert.Equal(t, []string{"seed_001", "seed_002", "seed_003"}, called)
}

package development

import (
	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder/config"
	"github.com/version-1/goooder/seeder"
)

type Seed struct {
	Seeders []config.Seeder
}

func NewSeeders(r *seeder.Renderer) []config.Seeder {
	seeders := []config.Seeder{
		&Seed_0000010_CreateUsers{r},
		&Seed_0000020_CreateProfiles{r},
	}

	return seeders
}

type Seed_0000010_CreateUsers struct {
	*seeder.Renderer
}

func (s *Seed_0000010_CreateUsers) Exec(tx *sqlx.Tx) error {
	return s.Renderer.ExecWithFilename(tx, "create_users.sql")
}

type Seed_0000020_CreateProfiles struct {
	*seeder.Renderer
}

func (s *Seed_0000020_CreateProfiles) Exec(tx *sqlx.Tx) error {
	return s.Renderer.ExecWithFilename(tx, "create_profiles.sql")
}

package development

import (
	"bytes"
	"html/template"

	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder/config"
)

type Seed struct {
	Seeders []config.Seeder
}

type Renderer struct {
	tmpl *template.Template
}

func (r Renderer) RenderTemplate(filename string) (string, error) {
	buf := new(bytes.Buffer)
	if err := r.tmpl.ExecuteTemplate(buf, filename, struct{}{}); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewRenderer() *Renderer {
	tmpl := template.Must(template.ParseGlob("db/seeders/development/templates/*.sql"))

	return &Renderer{
		tmpl: tmpl,
	}
}

func NewSeed() *Seed {
	tmpl := NewRenderer()

	return &Seed{
		Seeders: []config.Seeder{
			Seed_0000010_CreateUsers{r: tmpl},
			Seed_0000020_CreateFollows{r: tmpl},
		},
	}
}

type Seed_0000010_CreateUsers struct {
	r *Renderer
}

func (s Seed_0000010_CreateUsers) Exec(tx *sqlx.DB) error {
	query, err := s.r.RenderTemplate("create_users.sql")
	if err != nil {
		return err
	}

	tx.MustExec(query)
	return nil
}

type Seed_0000020_CreateFollows struct {
	r *Renderer
}

func (s Seed_0000020_CreateFollows) Exec(tx *sqlx.DB) error {
	query, err := s.r.RenderTemplate("create_follows.sql")
	if err != nil {
		return err
	}

	tx.MustExec(query)
	return nil
}

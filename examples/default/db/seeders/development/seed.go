package development

import (
	"bytes"
	"html/template"

	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder"
)

type Seed struct {
	Seeders []goooder.Seeder
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
		Seeders: []goooder.Seeder{
			Seed_0000010_CreateUsers{r: tmpl},
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

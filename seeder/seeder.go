package seeder

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder/config"
)

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer(pathGlob ...string) *Renderer {
	path := "db/seeders/development/templates/*.sql"
	if len(pathGlob) > 0 && len(pathGlob[0]) > 0 {
		path = pathGlob[0]
	}

	rootPath, err := LookUpFileInAncestors(".", "go.mod")
	if err != nil {
		panic(err)
	}

	glob := fmt.Sprintf("%s/%s", rootPath, path)

	tmpl := template.Must(template.ParseGlob(glob))

	return &Renderer{
		tmpl: tmpl,
	}
}

func (r *Renderer) RenderTemplate(filename string) (string, error) {
	buf := new(bytes.Buffer)
	if err := r.tmpl.ExecuteTemplate(buf, filename, struct{}{}); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (r *Renderer) ExecWithFilename(tx *sqlx.Tx, filename string) error {
	query, err := r.RenderTemplate(filename)
	if err != nil {
		return err
	}

	tx.MustExec(query)
	return nil
}

type Seed struct {
	r       *Renderer
	seeders []config.Seeder
}

func New(r *Renderer) *Seed {
	return &Seed{
		r: r,
	}
}

func LookUpFileInAncestors(path, filename string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if absPath == "/" {
		return "", fmt.Errorf("file not found in ancestors: %s", filename)
	}

	for _, e := range entries {
		if !e.IsDir() && e.Name() == filename {
			return absPath, nil
		}
	}

	parent := fmt.Sprintf("../%s", path)
	return LookUpFileInAncestors(parent, filename)
}

package core

import (
	"html/template"
	"path/filepath"
	"io/ioutil"
	"net/http"
	"strings"
	"errors"
)


type TemplateLoader struct {
	basePath 		string
	templates		map[string]*template.Template
}

func NewTemplateLoader(basePath string) (*TemplateLoader, error) {
	tl := TemplateLoader{}
	tl.basePath = basePath
	tl.templates = make(map[string]*template.Template)

	err := tl.initTemplates()
	if err != nil {
		return nil, err
	}

	return &tl, nil
}

func (tl *TemplateLoader) RenderTemplate(name string, w http.ResponseWriter, data interface{}) error {
	tmpl, ok := tl.templates[name]
	if !ok {
		return errors.New("The template " + name + " was not found.")
	}

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}

	return nil
}

func (tl *TemplateLoader) initTemplates() error {
	includeFiles, err := filepath.Glob(filepath.Join(tl.basePath, "include", "*.tmpl"))
	if err != nil {
		return err
	}

	// Prepare include templates here and clone them later.
	includeTmpl, err := template.ParseFiles(includeFiles...)
	if err != nil {
		return err
	}

	err = tl.loadTemplates("member", includeTmpl)
	if err != nil {
		return err
	}

	return nil
}

func (tl *TemplateLoader) loadTemplates(searchPath string, includeTmpl *template.Template) error {
	tmplFiles, err := filepath.Glob(filepath.Join(tl.basePath, searchPath, "*.tmpl"))
	if err != nil {
		return err
	}

	for _, tf := range tmplFiles {
		tfBytes, err := ioutil.ReadFile(tf)
		if err != nil {
			return err
		}
		tfStr := string(tfBytes)

		tn, err := filepath.Rel(tl.basePath, tf)
		if err != nil {
			return err
		}
		tn = strings.TrimSuffix(tn, ".tmpl")

		// We clone the included templates here to keep each namespace fresh.
		it, err := includeTmpl.Clone()
		if err != nil {
			return err
		}

		t, err := it.New(tn).Parse(tfStr)
		if err != nil {
			return err
		}

		tl.templates[tn] = t
	}

	return nil
}
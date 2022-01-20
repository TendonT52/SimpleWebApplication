package handlers

import (
	"net/http"

	"github.com/TendonT52/SimpleWebApplication/pkg/config"
	"github.com/TendonT52/SimpleWebApplication/pkg/models"
	"github.com/TendonT52/SimpleWebApplication/pkg/render"
)
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}

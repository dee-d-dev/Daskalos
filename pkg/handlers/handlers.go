package handlers

import (
	"net/http"

	"github.com/dee-d-dev/my_app/pkg/config"
	"github.com/dee-d-dev/my_app/pkg/models"
	"github.com/dee-d-dev/my_app/pkg/render"
)

//the repo used by the hnadlers
var Repo *Repository

//set repository type
type Repository struct {
	App *config.AppConfig
}

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//new handler
func NewHandlers(r *Repository) {
	Repo = r
}

//home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello there"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

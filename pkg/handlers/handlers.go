package handlers

import (
	"github.com/chirag3003/go-web/pkg/config"
	"github.com/chirag3003/go-web/pkg/models"
	"github.com/chirag3003/go-web/pkg/render"
	"log"
	"net/http"
)

var Repo *Repository

// TemplateData holds data send from handlers to templates

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world"
	render.RenderTemplate(res, "index", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	log.Println("about")
	res.Write([]byte(`{"page":"About"}`))
}

package handlers

import (
	"basic-server/pkg/config"
	"basic-server/pkg/models"
	"basic-server/pkg/render"
	"net/http"
)

// **REPO PATTERN
// hold in memory
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//create new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// set value to Repo
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	remoteip := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteip)

	stringMap := map[string]string{}
	stringMap["content"] = "string content works"
	stringMap["remote_ip"] = repo.App.Session.GetString(r.Context(), "remote_ip")
	td := models.TemplateData{StringMap: stringMap}

	render.RenderTemplate2(w, "home.page.html", &td)
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteip := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringmap := map[string]string{"remote_ip": remoteip}

	render.RenderTemplate2(w, "about.page.html", &models.TemplateData{StringMap: stringmap})
}

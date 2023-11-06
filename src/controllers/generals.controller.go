package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
)

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, r, "generals.page.tmpl", &config.TemplateData{})
}

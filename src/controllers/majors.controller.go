package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
)

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, r, "majors.page.tmpl", &config.TemplateData{})
}

package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/pkg/models"
	"github.com/yasniel1408/bookings/pkg/render"
)

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

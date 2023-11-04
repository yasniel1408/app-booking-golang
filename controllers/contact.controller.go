package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/pkg/models"
	"github.com/yasniel1408/bookings/pkg/render"
)

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

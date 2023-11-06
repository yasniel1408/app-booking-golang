package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
)

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	config.RenderTemplate(w, r, "contact.page.tmpl", &config.TemplateData{})
}

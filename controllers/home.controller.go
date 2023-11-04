package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/pkg/models"
	"github.com/yasniel1408/bookings/pkg/render"
)

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

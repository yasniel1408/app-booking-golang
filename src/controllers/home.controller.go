package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
)

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	config.RenderTemplate(w, r, "home.page.tmpl", &config.TemplateData{})
}

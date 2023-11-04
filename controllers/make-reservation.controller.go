package controllers

import (
	"net/http"

	"github.com/yasniel1408/bookings/pkg/models"
	"github.com/yasniel1408/bookings/pkg/render"
)

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

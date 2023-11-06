package controllers

import (
	"log"
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
	"github.com/yasniel1408/bookings/src/models"
)

// Reservation renders the reservation page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation") // clear reservation from session

	data := make(map[string]interface{})
	data["reservation"] = reservation

	config.RenderTemplate(w, r, "reservation-summary.page.tmpl", &config.TemplateData{
		Data: data,
	})
}

package controllers

import (
	"log"
	"net/http"

	"github.com/yasniel1408/bookings/src/config"
	"github.com/yasniel1408/bookings/src/config/forms"
	"github.com/yasniel1408/bookings/src/models"
)

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	config.RenderTemplate(w, r, "make-reservation.page.tmpl", &config.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handlers the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		config.RenderTemplate(w, r, "make-reservation.page.tmpl", &config.TemplateData{
			Form: form,
			Data: data,
		})
	}

	// save to session the reservation
	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

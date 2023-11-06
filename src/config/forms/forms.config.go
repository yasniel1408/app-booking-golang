package forms

import (
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors Errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		Errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if value == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Required checks for required fields
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// MinLength checks for strings of a minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, "This field must be at least "+string(length)+" characters long")
		return false
	}

	return true
}

// IsEmail checks for valid email addresses
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	mux := routes()

	switch v := mux.(type) {
	case *chi.Mux:
	// do nothing; test passed
	default:
		t.Error(fmt.Printf("type is not *chi.Mux, but is %T", v))
	}
}

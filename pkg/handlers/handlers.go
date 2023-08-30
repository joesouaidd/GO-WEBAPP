package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "./templates/home.page.tmpl")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "./templates/about.page.tmpl")
}

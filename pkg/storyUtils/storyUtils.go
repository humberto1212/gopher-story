package storyutils

import (
	"html/template"
	"net/http"

	"github.com/humberto1212/gopher-story/pkg/utils"
)

func IntroHandler(w http.ResponseWriter, r *http.Request) {

	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["intro"])

}

func NewYorkHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["new-york"])
}

func DebateHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["debate"])
}
func SeanKellyHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["sean-kelly"])
}
func MarkBatesHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["mark-bates"])
}
func DenverHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["denver"])
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	story := utils.ParseJSON()

	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, story["home"])
}

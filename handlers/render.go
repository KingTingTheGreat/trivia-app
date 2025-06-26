package handlers

import (
	"html/template"
	"io"
	"net/http"
	"strings"
	"trivia-app/dlog"
	"trivia-app/views"
)

type Play struct {
	Ready bool
}

func RenderTemplate(w http.ResponseWriter, filename string, data any) {
	// parse layout
	tmpl := template.Must(template.New("layout").Parse(views.Layout))
	// parse pages
	tmpl = template.Must(tmpl.ParseFS(views.Pages, "pages/"+filename))
	// parse components
	tmpl = template.Must(tmpl.ParseFS(views.Components, "components/*.html"))

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		dlog.DLog("error rendering page")
		dlog.DLog(err)
		return
	}
}

func RenderComponent(w io.Writer, filename string, data any) {
	// parse component
	tmpl := template.Must(template.ParseFS(views.Components, "components/"+filename))

	err := tmpl.ExecuteTemplate(w, strings.TrimRight(filename, ".html"), data)
	if err != nil {
		dlog.DLog("error rendering component")
		dlog.DLog(err)
		return
	}
}

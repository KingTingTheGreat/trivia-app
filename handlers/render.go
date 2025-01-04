package handlers

import (
	"html/template"
	"io"
	"strings"
	"trivia-app/dlog"
)

type Count struct {
	Count int
}

type Play struct {
	Ready bool
}

func RenderTemplate(w io.Writer, filename string, data interface{}) {
	// load layout and page files
	tmpl := template.Must(template.New("layout").ParseFiles("views/layout.html", "views/pages/"+filename))
	// load component files
	tmpl = template.Must(tmpl.ParseGlob("views/components/*.html"))

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		dlog.DLog("error rendering page")
		dlog.DLog(err)
		return
	}
}

func RenderComponent(w io.Writer, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("views/components/" + filename))

	err := tmpl.ExecuteTemplate(w, strings.TrimRight(filename, ".html"), data)
	if err != nil {
		dlog.DLog("error rendering component")
		dlog.DLog(err)
		return
	}
}

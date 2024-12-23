package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/migopp/gocards/internal/debug"
)

func serveTmpl(w http.ResponseWriter, tmplName string, dynContent DynContent) {
	// Templates live in `/web/templates`
	tmplPath := filepath.Join("web", "templates", tmplName)
	debug.Printf("| Looking for template `%s` @ %v\n", tmplName, tmplPath)

	// Parse the template with name `tname`
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		errStr := fmt.Sprintf("ERROR LOADING TEMPLATE %s [%v]", tmplName, err)
		http.Error(w, errStr, http.StatusInternalServerError)
	}

	// Actually load the template by writing to the response
	//
	// Looks like we can send dynamic data as well:
	// https://pkg.go.dev/html/template#Template.Execute
	//
	// I don't really know how this works for now so it's a later problem
	err = tmpl.Execute(w, dynContent)
	if err != nil {
		errStr := fmt.Sprintf("ERROR EXECUTING TEMPLATE %s [%v]", tmplName, err)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}

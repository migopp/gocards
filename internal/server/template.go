package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func servePage(w http.ResponseWriter, tmplName string) error {
	// Templates live in `/web/templates`
	tmplPath := filepath.Join("web", "templates", tmplName)

	// Parse the template with name `tname`
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		errStr := fmt.Sprintf("ERROR LOADING TEMPLATE %s [%v]", tmplName, err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return fmt.Errorf(errStr)
	}

	// Actually load the template by writing to the response
	//
	// Looks like we can send dynamic data as well:
	// https://pkg.go.dev/html/template#Template.Execute
	//
	// I don't really know how this works for now so it's a later problem
	err = tmpl.Execute(w, nil)
	if err != nil {
		errStr := fmt.Sprintf("ERROR EXECUTING TEMPLATE %s [%v]", tmplName, err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return fmt.Errorf(errStr)
	}

	return nil
}

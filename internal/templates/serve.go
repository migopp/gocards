package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

func ServeTemplate(w http.ResponseWriter, t string, c any) {
	// Parse `t` as an HTML template
	tmpl, err := template.New("tmpl").Parse(t)
	if err != nil {
		errStr := fmt.Sprintf("ERROR LOADING TEMPLATE [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
	}

	// Write our HTML to the DOM
	err = tmpl.Execute(w, c)
	if err != nil {
		errStr := fmt.Sprintf("ERROR EXECUTING TEMPLATE [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}

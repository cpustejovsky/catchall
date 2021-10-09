package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) updateDelivered(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get(":domain_name")

	err := app.domains.UpdateDelivered(domain)
	if err != nil {
		fmt.Println(err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Successfully updated number of delivered emails for %v", domain)
}

func (app *application) updateBounced(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get(":domain_name")

	err := app.domains.UpdateBounced(domain)
	if err != nil {
		fmt.Println(err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Successfully updated number of bounced emails for %v", domain)
}

func (app *application) checkStatus(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get(":domain_name")

	status, err := app.domains.CheckStatus(domain)
	if err != nil {
		fmt.Println(err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "Domain %v is status %v", domain, status)
}

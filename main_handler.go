package main

import "net/http"

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/test
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplete(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/edit
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplete(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /save/
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

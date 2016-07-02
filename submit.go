// Copyright 2015 Robert S. Gerus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/monoculum/formam"
)

var _ = path.Join

type entry struct {
	Text     string    `form:"text"`
	Who      string    `form:"who"`
	Location string    `form:"location"`
	When     time.Time `form:"when"`
	From     time.Time `form:"from"`
	To       time.Time `form:"to"`
}

func submit(w http.ResponseWriter, r *http.Request) {
	layout := path.Join("templates", "layout.html")
	form := path.Join("templates", "submit.html")

	tmpl, err := template.ParseFiles(layout, form)
	if err != nil {
		log.Println("error paring templates", err)
		http.Error(w, http.StatusText(500), 500)
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func timestampDecoder(vals []string) (interface{}, error) {
	return time.Parse("2006-01-02 15:04", vals[0])
}

func submit_form(w http.ResponseWriter, r *http.Request) {
	var e entry
	layout := path.Join("templates", "layout.html")
	form := path.Join("templates", "submit.html")

	tmpl, _ := template.ParseFiles(layout, form)
	r.ParseForm()

	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "form"})

	dec.RegisterCustomType(timestampDecoder, []interface{}{time.Time{}}, []interface{}{&e.When, &e.From, &e.To})

	if err := dec.Decode(r.Form, &e); err != nil {
		log.Println("error decoding form", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", e)

	log.Printf("entry: %+v\n", e)
}

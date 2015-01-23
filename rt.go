// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

type Translation struct {
	Title string
	Source  []byte
	Target  []byte
}

func (tr *Translation) save() error {
	filename := tr.Title + "_t.txt"
	return ioutil.WriteFile(filename, tr.Target, 0600)
}

//This is a comment
func loadPage(title string) (*Translation, error) {
	filename := title + ".txt"
	source, err := ioutil.ReadFile(filename)
	filename_target := title + "_t.txt"
	target, err := ioutil.ReadFile(filename_target)
	if err != nil {
		return nil, err
	}
	return &Translation{Title: title, Source: source, Target: target}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	tr, _ := loadPage(title)
	renderTemplate(w, "main", tr)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	target := r.FormValue("target")
	tr := &Translation{Title: title, Source: []byte(target), Target: []byte(target)}
	err := tr.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("main.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, tr *Translation) {
	err := templates.ExecuteTemplate(w, tmpl+".html", tr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}


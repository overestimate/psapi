package main

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
)

func declareHandlers() {
	http.HandleFunc("/docs", fetchFile)
	http.HandleFunc("/droptime/", namemcSearchPage)
	http.HandleFunc("/", namemcReplier)
}

func getDocs(w http.ResponseWriter, r *http.Request) {
	p := "htdocs/index.htm"
	t, e := readFile(p)
	if e != nil || t == nil {
		w.WriteHeader(404)
		w.Write([]byte("<h1>PAGE NOT FOUND</h1><br>Something's wrong with the link, if this is a valid page contact me @ t.me/tringram"))
		return
	}
	w.Header().Set("content-type", mime.TypeByExtension(filepath.Ext(p)))
	fmt.Fprintf(w, "%s", *t)
}

func fetchFile(w http.ResponseWriter, r *http.Request) {
	p := "htdocs" + r.URL.Path
	if p == "htdocs/" {
		p = "htdocs/index.htm"
	}
	t, e := readFile(p)
	if e != nil || t == nil {
		w.WriteHeader(404)
		w.Write([]byte("<h1>PAGE NOT FOUND</h1><br>Something's wrong with the link, if this is a valid page contact me @ t.me/tringram"))
		return
	}
	w.Header().Set("content-type", mime.TypeByExtension(filepath.Ext(p)))
	fmt.Fprintf(w, "%s", *t)
}

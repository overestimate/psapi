package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
This file should contain all REST API endpoint functions.
Handlers should be put in handlers.go.
*/

/*
 ooooo namemc!
*/

func namemcReplier(w http.ResponseWriter, r *http.Request) {
	// this should mirror a reply from namemc in status code, content, and content type

	var req *http.Request
	query := ""
	if query != r.URL.RawQuery {
		query = "?" + r.URL.RawQuery
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.namemc.com%s%s", r.URL.Path, query), nil)

	client := &http.Client{}

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("request not generated, cannot continue"))
		return
	}
	req.Host = "namemc.com"
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("no reply from namemc, cant continue"))
		return
	}
	m, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("cant read reply from namemc"))
		return
	}
	w.WriteHeader(res.StatusCode)
	w.Write(m)
	fmt.Println(m)
}

func namemcSearchPage(w http.ResponseWriter, r *http.Request) {
	namespl := strings.Split(r.URL.Path, "/")
	name := namespl[len(namespl)-1]
	if name == "" {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"err\": \"no name provided\"}"))
		return
	}
	fmt.Println(name)
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.namemc.com/search?q=%s", name), nil)

	if err != nil {
		resp := "{\"err\": \"failed to generate req for lookup.\"}"
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
		return
	}
	req.Host = "namemc.com"
	res, err := client.Do(req)
	if err != nil {
		resp := "{\"err\": \"lookup request recieved error on reply.\"}"
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
		return
	}
	fmt.Printf("successful request made for username %s\n", name)
	s, err := io.ReadAll(res.Body)
	if err != nil {
		resp := "{\"err\": \"body is missing, cannot continue.\"}"
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resp))
		return
	}
	searchesStr := strings.Split(string(s)[strings.Index(string(s), searchesPrecursorString)+len(searchesPrecursorString):], " /")[0]
	statusStr := strings.Split(string(s)[strings.Index(string(s), statusPrecursorString)+len(statusPrecursorString):], "</div>")[0]
	if statusStr == "Available Later" {
		timestampStr := strings.Split(string(s)[strings.Index(string(s), timestampPrecursorString)+len(timestampPrecursorString):], "\">")[0]
		unixTimeObj, err := time.Parse(time.RFC3339, timestampStr)
		if err != nil {
			fmt.Println("err=" + err.Error())
			resp := "{\"err\": \"timestamp failed conversion.\"}"
			w.WriteHeader(500)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(resp))
			return
		}
		unixI64 := unixTimeObj.UnixNano() / int64(time.Millisecond)
		rs := "{\n  \"unix\": " + strconv.FormatInt(unixI64, 10) + ",\n  \"droptime\": \"" + timestampStr + "\",\n  \"username\": \"" + name + "\",\n  \"searches\": " + searchesStr + ",\n  \"status\": \"" + statusStr + "\"\n}"
		w.Write([]byte(rs))
		return
	}
	rs := "{\n  \"username\": \"" + name + "\",\n  \"searches\": " + searchesStr + ",\n  \"status\": \"" + statusStr + "\"\n}"
	w.Write([]byte(rs))
	fmt.Println(s)
}

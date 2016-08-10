package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSubmit(t *testing.T) {
	name := "Karl"
	password := "proletariat"
	v := url.Values{}
	v.Set("left_displayname", name)
	v.Set("employee_passwd", password)
	v.Set("left_inout", "in")
	v.Set("left_notes", "")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if v.Encode() != string(body) {
			t.Fatalf("expected %v got %v", v.Encode(), string(body))
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	if err := submit(&http.Client{}, server.URL, v); err != nil {
		t.Fatal(err)
	}
}

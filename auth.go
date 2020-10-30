package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)


func checkBasicAuth(w http.ResponseWriter, r *http.Request) (bool) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", 403)
		return false
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validate(pair[0], pair[1]) {
		http.Error(w, "authorization failed", 403)
		return false
	}

	return true
}

func validate(username, password string) bool {
	if username == credentials["username"] && password == credentials["password"] {
		return true
	}
	return false
}


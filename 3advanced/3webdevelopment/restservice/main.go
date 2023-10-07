package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/weapon":
		weapon(w, r)
	default:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello Ninja!")
	}
}

var ninjaWeapons = map[string]string{
	"ninjaStar":  "Baby Ninja Star - Level 1, Damage 1",
	"ninjaSword": "Baby Ninja Sword - Level 1, Damage 1",
}

func weapon(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		badRequest(w, r)
		return
	}
	weapon, err := getWeapon(w, r)
	if err != nil {
		badRequest(w, r)
		return
	}
	writeResponse(weapon, w, r)
}

func writeResponse(weapon string, w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Add("Content-Type", "application/json")
		weapon = fmt.Sprintf("{\"weapon\":\"%s\"}", weapon)
	case "text/html":
		w.Header().Add("Content-Type", "text/html")
		weapon = fmt.Sprintf("<h1>%s</h1>", weapon)
	default:
		w.Header().Add("Content-Type", "text/plain")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, weapon)
}

func getWeapon(w http.ResponseWriter, r *http.Request) (string, error) {
	weaponName, err := getWeaponName(w, r)
	if err != nil {
		return "", errors.New("failed to get weapon name")
	}
	weapon := ninjaWeapons[weaponName]
	return weapon, nil
}

func getWeaponName(w http.ResponseWriter, r *http.Request) (string, error) {
	weaponName := r.FormValue("name")
	if len(weaponName) != 0 {
		return weaponName, nil
	}

	type WeaponRequestBody struct {
		Name string
	}
	var body WeaponRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return "", err
	}
	return body.Name, nil
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Bad request!")
}

func main() {
	// REST service

	// Establish a service
	var handler http.ServeMux
	handler.HandleFunc("/", handleFunc)
	server := http.Server{
		Addr:         "", // localhost:80 (host:port)
		Handler:      &handler,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	server.ListenAndServe()
}

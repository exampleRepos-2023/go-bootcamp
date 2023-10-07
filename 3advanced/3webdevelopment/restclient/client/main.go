package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// REST main

	// Plain old URL - http://localhost
	response, err := http.Get("http://localhost/weapon?name=ninjaStar")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Header.Get("Content-Type"))
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}

	// Request method, header, body...
	request, err := http.NewRequest(
		"GET",
		"http://localhost/weapon",
		strings.NewReader(`{"name": "ninjaStar"}`))
	request.Header.Set("Accept", "application/json")
	client := http.Client{}
	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Header.Get("Content-Type"))
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

package main

import (
	"fmt"
	"net/http"
)

func pingDomain(domain string) string {
	resp, err := http.Get("http://" + domain)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	return resp.Status
}

func main() {
	result := pingDomain("example.com")
	fmt.Println(result)
}

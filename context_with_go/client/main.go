package main

import (
	"io"
	"net/http"
	"os"

	"log"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	io.Copy(os.Stdout, resp.Body)
}

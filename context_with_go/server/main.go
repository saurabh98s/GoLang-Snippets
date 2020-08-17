package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("started handler")
	defer log.Println("ended handler")

	select {
	case <-time.After(5*time.Second):
		fmt.Fprintln(w, "Hello ")
	case <-ctx.Done():
		err:=ctx.Err()
		log.Print(err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	
}

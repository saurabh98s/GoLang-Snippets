package main

import (
	"fmt"
	log "context_with_go/logging"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx,"started handler")
	defer log.Println(ctx,"ended handler")
	// it wont send because the http package doesnt allow it
	// for cancellation we have a ack from tcp that conn has dropped
	// no such thing for sending data across server
	// write code for that
	fmt.Printf("Value of foo %v",ctx.Value("foo"))
	select {
	case <-time.After(5*time.Second):
		fmt.Fprintln(w, "Hello ")
	case <-ctx.Done():
		err:=ctx.Err()
		log.Println(ctx,err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	
}

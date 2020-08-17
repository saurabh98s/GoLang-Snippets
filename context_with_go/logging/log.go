package logging

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const requestIDKey = 43

// Println sends the id with the message
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("[ERROR] cant find request id in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

// Decorate returns a handler
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		ctx:=r.Context()
		id:=rand.Int63()
		ctx= context.WithValue(ctx,requestIDKey,id)
		f(w,r.WithContext(ctx))
	}
}

package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"log"
)

func main() {
	ctx := context.Background()
	ctx,cancel :=context.WithTimeout(ctx,time.Second)
	defer cancel()
	ctx= context.WithValue(ctx,"foo","bar")
	req,err:=http.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	if err != nil {
		log.Fatal(err)
	}
	req=req.WithContext(ctx)
	resp,err:= http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	io.Copy(os.Stdout, resp.Body)
}

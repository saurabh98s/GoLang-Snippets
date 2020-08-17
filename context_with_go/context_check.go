package main

import (
	"context"
	"fmt"
	"time"
)


func sleepAndTalk(ctx context.Context,timer time.Duration,text string) {
	time.Sleep(timer)
	fmt.Println(text)
}

func main() {
	ctx:=context.Background()
	sleepAndTalk(ctx,5*time.Second,"jello")
}
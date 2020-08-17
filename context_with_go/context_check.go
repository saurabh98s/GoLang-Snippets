package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	sleepAndTalk(ctx, 4*time.Second, "yoyoyo")
}

func sleepAndTalk(ctx context.Context, timer time.Duration, text string) {
	select {
	case <-time.After(timer):
		fmt.Println(text)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}

}

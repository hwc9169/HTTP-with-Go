package main

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// 1초당 상한 횟수
	RateLimit := 10
	// 토큰 최대 보유 수
	BucketSize := 10
	ctx := context.Background()
	e := rate.Every(time.Second / RateLimit)
	l := rate.NewLimiter(e, BucketSize)

	for _, task := range tasks {
		err := l.Wait(ctx)
		if err != nil {
			panic(err)
		}
	}
}

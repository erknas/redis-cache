package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	var (
		rCache = NewRedisCache(rClient)
		store  = NewStorage(rCache)
		ctx    = context.Background()
	)

	for i := 1; i <= 5; i++ {
		val, err := store.Get(ctx, "1")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", val)

		time.Sleep(time.Second * 3)
	}
}

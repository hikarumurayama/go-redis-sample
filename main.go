package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Data struct {
	key   string
	value string
}

func main() {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	d := Data{
		key:   "key1",
		value: "value1",
	}

	dd := Data{
		key:   "key1",
		value: "value100",
	}

	var ctx = context.Background()

	go func() {
		if err := Set(c, d, ctx); err != nil {
			fmt.Printf("key:%s value:%s %v\n", d.key, d.value, err)
		}
	}()

	go func() {
		if err := Set(c, dd, ctx); err != nil {
			fmt.Printf("key:%s value:%s %v\n", dd.key, dd.value, err)
		}
	}()

	time.Sleep(3 * time.Second)
}

func Set(c *redis.Client, d Data, ctx context.Context) error {
	// 値が存在していないときのみ値をセットする
	ok, err := c.SetNX(ctx, d.key, d.value, 20*time.Second).Result()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to set")
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
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

	var ctx = context.Background()

	// 登録
	if err := c.Set(ctx, d.key, d.value, 0).Err(); err != nil {
		panic(err)
	}

	// 取得
	val, err := c.Get(ctx, d.key).Result()
	switch {
	case err == redis.Nil:
		panic("key does not exist")
	case err != nil:
		panic(err)
	case val == "":
		panic("value is empty")
	}

	fmt.Println(d.key, val)
}

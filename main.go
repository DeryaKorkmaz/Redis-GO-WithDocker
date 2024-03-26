package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	ctx := context.Background()

	// Redis istemcisini oluştur
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380", // Docker Compose'da belirtilen servis adı
		Password: "",               // Redis parolası (varsayılan olarak boş)
		DB:       0,                // Redis veritabanı
	})

	json, err := json.Marshal(Employee{Name: "Derya", Address: "Turkey"})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, "staff", json, 0).Err()
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	val, err := client.Get(ctx, "staff").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}

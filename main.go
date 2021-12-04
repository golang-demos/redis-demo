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

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	jsonData, err := json.Marshal(Employee{Name: "Gordon", Address: "San Jose"})
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set(ctx, "add123", jsonData, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get(ctx, "add123").Result()
	if err != nil {
		fmt.Println(err)
	}

	emp1 := new(Employee)

	json.Unmarshal([]byte(val), &emp1)

	fmt.Println(emp1.Name)
	fmt.Println(emp1.Address)
}

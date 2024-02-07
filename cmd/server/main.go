package main

import (
	"easyshop-api/configs"

	_ "github.com/lib/pq"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}

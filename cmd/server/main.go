package main

import "easyshop-api/configs"

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}

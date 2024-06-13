package main

import "APIS/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}

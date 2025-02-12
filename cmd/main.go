package main

import (
	"price_checker/api"
	"price_checker/storage"
)

func main() {
	storage.ConnectDB()
	api.ListingsLatest()
}

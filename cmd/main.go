package main

import (
	"price_checker/CMCservice"
	"price_checker/storage"
)

func main() {
	db := storage.ConnectDB()
	repo := storage.NewRequestBDRepository(db)

	CMCservice.ListingsLatest(repo)
}

package main

import (
	"guest-ledger-book-api/db"
	"guest-ledger-book-api/router"
)

func main() {
	db.Db()
	router.Router()
}

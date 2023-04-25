package main

import (
	"Moddormy_backend/loaders/fiber"
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/storage"
)

func main() {
	mysql.Init()
	fiber.Init()
	storage.Init()
}

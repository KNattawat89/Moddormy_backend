package main

import (
	"Moddormy_backend/loaders/fiber"
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/storage"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	mysql.Init()
	fiber.Init()
	storage.Init()
}

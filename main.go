package main

import (
	"Moddormy_backend/loaders/fiber"
	"Moddormy_backend/loaders/mysql"
)

func main() {
	mysql.Init()
	fiber.Init()
}

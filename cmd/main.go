package main

import (
	"web-cache/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	app.Run()
}

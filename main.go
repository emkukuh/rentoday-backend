package main

import (
	"fmt"
	"os"

	"rentoday.id/app"
)

func main() {
	fmt.Println(os.Getenv("DB_HOST"))
	app.Run()
}
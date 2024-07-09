package main

import (
	"fmt"

	"example.com/tracker/internal/app"
)

func main() {
	app := app.NewApp()
	if err := app.SetupConfig("../config/.env"); err != nil {
		fmt.Println(err)
	}
	app.Run()
}

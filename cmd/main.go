package main

import (
	"github.com/SerjLeo/finance_bot/internal/app"
	"os"
)

func main() {
	err := app.Run()
	if err != nil {
		os.Exit(2)
	}
}

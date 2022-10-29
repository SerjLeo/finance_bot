package main

import (
	"fmt"
	"github.com/SerjLeo/finance_bot/internal/app"
	"os"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

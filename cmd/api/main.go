package main

import (
	"fmt"

	"github.com/gef3dx/api_workinghub/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}
}

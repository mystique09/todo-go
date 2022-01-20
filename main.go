package main

import (
	app "server-go/init"
)

func main() {
	app := app.App{}
	app.Init()
	app.Run()
}

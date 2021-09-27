package main

import (
	config "berbagi/config"
	routes "berbagi/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

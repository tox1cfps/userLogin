package main

import (
	"userLogin/config"
	"userLogin/internal/ui"
)

func main() {
	config.Conect()
	ui.Menu()
}

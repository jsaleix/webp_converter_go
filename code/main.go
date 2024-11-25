package main

import (
	"webp_converter/config"
	"webp_converter/converter"
)

func main() {
	config.Init()
	converter.Run()
}

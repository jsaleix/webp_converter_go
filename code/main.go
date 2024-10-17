package main

import (
	"webp_converter/config"
	"webp_converter/converter"
)

func main() {
	config.Init()

	options := config.GetOptions()

	if options.HelpValue {
		converter.GetHelp()
	} else {
		converter.Run()
	}
}

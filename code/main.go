package main

import (
	"flag"
	"webp_converter/converter"
)

func main() {
	helpCmd := flag.Bool("help", false, "Show the folders")

	flag.Parse()

	if *helpCmd {
		converter.GetHelp()
	} else {
		converter.Run()
	}
}

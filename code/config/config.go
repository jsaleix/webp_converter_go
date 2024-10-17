package config

import "flag"

type CLIOptions struct {
	HelpValue bool
	DevMode   bool
}

var options = CLIOptions{
	HelpValue: false,
	DevMode:   false,
}

func Init() {
	helpOpt := flag.Bool("help", false, "Show the folders")
	devOpt := flag.Bool("dev", false, "For dev purpose it hardcodes the executable path")

	flag.Parse()

	options.HelpValue = *helpOpt
	options.DevMode = *devOpt
}

func GetOptions() CLIOptions {
	return options
}

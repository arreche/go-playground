package cmd

import (
	"flag"
	"log"
	"os"
)

// Execute cmd
func Execute() {
	command := flag.NewFlagSet("test", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("missing command")
	}

	switch flag.Arg(0) {
	case "cmd1":
		log.Printf("running %s", flag.Arg(0))
	case "cmd2":
		mode := command.String("m", "soft", "mode")
		command.Parse(os.Args[2:])
		log.Printf("running %s in %s mode", *mode, flag.Arg(0))
	default:
		log.Fatal("unsupported command")
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tristnaja/Ucheck/cmd"
)

func main() {
	log.SetPrefix("ucheck: ")
	log.SetFlags(0)

	cfgRoot, err := os.UserConfigDir()

	if err != nil {
		log.Fatal(err)
	}

	dbDir := filepath.Join(cfgRoot, "ucheck")
	filePath := filepath.Join(dbDir, "db.json")

	if len(os.Args) < 1 && os.Args[1] == "run" {
		fmt.Fprintf(os.Stderr, "usage: ucheck run")
		log.Fatal("error occured: ucheck run does not need any arguments")
	} else if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ucheck <cmd> <args>")
		log.Fatal("insufficient args: add proper args")
	}

	switch os.Args[1] {
	case "run":
		cmd.RunExecute(filePath)
	case "add":
		err = cmd.RunAdd(os.Args[2:], filePath)
	case "list":
		err = cmd.RunList(filePath)
	default:
		log.Fatal("unknown command, usable: run, add, delete, deleteAll")
	}

	if err != nil {
		log.Fatal(err)
	}
}

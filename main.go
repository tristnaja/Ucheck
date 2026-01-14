package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tristnaja/Ucheck/cmd"
)

func main() {
	log.SetPrefix("ucheck: ")
	log.SetFlags(0)

	urls := []string{
		"https://google.com", "https://github.com",
		"https://go.dev", "https://stackoverflow.com",
		"https://nonexistent.site.example", // Err Test
	}

	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: ucheck <cmd> <args>")
		log.Fatal("parsing args: arguments not enough")
	}

	switch os.Args[1] {
	case "run":
		cmd.RunExecute(urls)
	default:
		log.Fatal("unknown command, usable: run, add, delete, deleteAll")
	}
}

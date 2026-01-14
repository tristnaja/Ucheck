package cmd

import (
	"flag"
	"fmt"

	"github.com/tristnaja/Ucheck/internal"
)

func RunAdd(args []string, filePath string) error {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)
	var url string

	cmd.StringVar(&url, "url", "", "New URL for Queue")
	cmd.StringVar(&url, "u", "", "New URL for Queue (shorthand)")

	err := cmd.Parse(args)

	if err != nil {
		return fmt.Errorf("parsing arguments: %w", err)
	}

	if url == "" {
		cmd.Usage()
		return fmt.Errorf("provide a new link")
	}

	err = internal.AddURL(url, filePath)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully add new URL: %v\n", url)

	return nil
}

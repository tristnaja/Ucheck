package cmd

import (
	"fmt"

	"github.com/tristnaja/Ucheck/internal"
)

func RunList(filePath string) error {
	db, err := internal.ReadDatabase(filePath)

	if err != nil {
		return err
	}

	fmt.Println("Here are the list of your queue:")
	fmt.Printf("| Size | %d\n", db.Size)
	for index, url := range db.URLs {
		fmt.Printf("%d. %v\n", (index + 1), url)
	}

	return nil
}

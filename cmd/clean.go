package cmd

import "github.com/tristnaja/Ucheck/internal"

func RunClean(filePath string) error {
	err := internal.CleanUp(filePath)

	if err != nil {
		return err
	}

	return nil
}

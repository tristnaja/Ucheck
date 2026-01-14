package cmd

import (
	"github.com/tristnaja/Ucheck/internal"
)

func RunExecute(filePath string) error {
	err := internal.RunCheck(filePath)

	if err != nil {
		return err
	}

	return nil
}

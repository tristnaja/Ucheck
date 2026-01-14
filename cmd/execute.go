package cmd

import (
	"github.com/tristnaja/Ucheck/internal"
)

func RunExecute(filePath string) {
	internal.RunCheck(filePath)
}

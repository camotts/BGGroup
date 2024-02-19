package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
	Short: "bggroup",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

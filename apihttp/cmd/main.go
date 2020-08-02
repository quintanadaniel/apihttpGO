package main

import (
	"github.com/CodelyTV/test_project/apihttp/internal"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "Beers"}
	rootCmd.AddCommand(internal.InitBeer())
	rootCmd.Execute()

}

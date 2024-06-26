/*
Copyright © 2024 Kanstantsin Markelau
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/kmarkela/wiggumizer/internal/historyparser"
	"github.com/spf13/cobra"
)

const version = "v1.0.4"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wiggumizer",
	Short: "Web Traffic 4nalizer",
	Long:  `Web Traffic 4nalizer`,
	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Printf("Wiggumizer: %s\n", version)
			os.Exit(0)
		}
		cmd.Root().Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "V", false, "print version")
	// rootCmd.PersistentFlags().StringP("historyFile", "f", "", "path to history file")

	rootCmd.PersistentFlags().IntP("workers", "w", 5, "amount of workers")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

}

func getHistory(fname string) *historyparser.BrowseHistory {

	// read history file
	data, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//parser history
	var bh = &historyparser.BrowseHistory{}
	historyparser.ParseHistory(&data, bh)

	// define scope
	sh := GetMUltipleChoices("Choose hosts in Scope:", bh.ListOfHosts.Keys())
	bh.FilterByHost(sh)

	return bh
}

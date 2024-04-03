/*
Copyright © 2024 Kanstantsin Markelau
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/spf13/cobra"
)

const version = "v0.2.0-alpha"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Wiggumize",
	Short: "Web Traffic 4nalizer",
	Long:  `Web Traffic 4nalizer`,
	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Printf("Wiggumizer: %s", version)
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

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/kmarkela/Wiggumizeng/cmd"
// 	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
// 	"github.com/kmarkela/Wiggumizeng/internal/scanner"
// 	"github.com/kmarkela/Wiggumizeng/internal/search"
// )

// const version = "0.1.1"

// func main() {

// 	// init cmd
// 	wgr, err := cmd.NewWiggumiser()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	// print version and exit
// 	if wgr.Params.Action == cmd.Version {
// 		fmt.Printf("Wiggumizer: %s", version)
// 		return
// 	}

// 	// print ASCII art
// 	wgr.Greet()

// 	// read history file
// 	data, err := os.ReadFile(wgr.Params.History)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	// parser history
// 	var bh = &historyparser.BrowseHistory{}
// 	historyparser.ParseHistory(&data, bh)

// 	// define scope
// 	sh := cmd.GetMUltipleChoices("Choose hosts in Scope:", bh.ListOfHosts.Keys())
// 	bh.FilterByHost(sh)

// 	// check action
// 	switch wgr.Params.Action {
// 	case cmd.Search:
// 		s := search.Searcher{}
// 		s.Search(bh, wgr.Params.Workers)
// 	case cmd.Scan:
// 		sc := scanner.NewScanner()
// 		sc.Scan(bh, wgr.Params.Output, wgr.Params.Workers)
// 	}
// }

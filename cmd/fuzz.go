/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/kmarkela/Wiggumizeng/internal/fuzz"
	"github.com/spf13/cobra"
)

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "fuzz all endpoint from history",
	Long:  `It allows to fuzz muptiple parameters over multiple endpoints`,
	Run: func(cmd *cobra.Command, args []string) {

		// get params
		fname, err := cmd.Flags().GetString("wordlist")
		if err != nil {
			log.Fatalln(err)
		}
		hfname, err := cmd.Flags().GetString("historyFile")
		if err != nil {
			log.Fatalln(err)
		}
		proxy, err := cmd.Flags().GetString("proxy")
		if err != nil {
			log.Fatalln(err)
		}
		workers, err := cmd.Flags().GetInt("workers")
		if err != nil {
			log.Fatalln(err)
		}
		maxReq, err := cmd.Flags().GetInt("maxReq")
		if err != nil {
			log.Fatalln(err)
		}
		headers, err := cmd.Flags().GetStringSlice("headers")
		if err != nil {
			log.Fatalln(err)
		}
		excludeEndpoint, err := cmd.Flags().GetStringSlice("excludeEndpoint")
		if err != nil {
			log.Fatalln(err)
		}
		excludeParam, err := cmd.Flags().GetStringSlice("excludeParam")
		if err != nil {
			log.Fatalln(err)
		}
		parameters, err := cmd.Flags().GetStringSlice("parameters")
		if err != nil {
			log.Fatalln(err)
		}
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatalln(err)
		}

		f, err := fuzz.New(workers, maxReq, headers, excludeEndpoint, excludeParam, parameters, fname, proxy, verbose)
		if err != nil {
			log.Fatalln(err)
		}

		greet()

		bh := getHistory(hfname)

		f.Run(bh)

	},
}

func init() {
	fuzzCmd.Flags().StringP("historyFile", "f", "", "path to history file")
	fuzzCmd.MarkFlagRequired("historyFile")

	fuzzCmd.Flags().StringP("wordlist", "l", "", "wordlits to fuzz")
	fuzzCmd.MarkFlagRequired("wordlist")
	fuzzCmd.Flags().StringP("proxy", "p", "", "proxy")
	fuzzCmd.Flags().IntP("maxReq", "m", 0, "max amount of requests per second")
	fuzzCmd.Flags().StringSlice("headers", []string{}, "ISN'T IMPLEMENTED YET. replace header if exists, add if it wasn't in original request")
	fuzzCmd.Flags().StringSlice("excludeEndpoint", []string{}, "ISN'T IMPLEMENTED YET. exclude specific endpoints from fuzz")
	fuzzCmd.Flags().StringSlice("excludeParam", []string{}, "ISN'T IMPLEMENTED YET. exclude specific parameters from fuzz")
	fuzzCmd.Flags().StringSlice("parameters", []string{}, "ISN'T IMPLEMENTED YET. fuzz only specified parameters")

	fuzzCmd.Flags().BoolP("verbose", "v", false, "ISN'T IMPLEMENTED YET. verbose")

	rootCmd.AddCommand(fuzzCmd)
}

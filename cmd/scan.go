package cmd

import (
	"github.com/kmarkela/Wiggumizeng/internal/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan analysis web history and run list of checks on Req\\Res body and headers",
	Long: `
List of checks

- LFI Checker: This module is searching for filenames in request parameters. Could be an indication of possible LFI
- Redirect Checker: This module is searching for Redirects
- Secret Checker: This module lokking for sensitive information, such as API keys
- SSRF Checker: This module is searching for URL in request parameters.
- Subdomain Checker: This module is searching for 404 messages form hosting platformas
- XML Checker: This module is searching for XML in request parameters
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// get params
		fname, _ := cmd.Flags().GetString("historyFile")
		workers, _ := cmd.Flags().GetInt("workers")
		output, _ := cmd.Flags().GetString("output")
		verbose, _ := cmd.Flags().GetBool("verbose")

		greet()

		bh := getHistory(fname)

		sc := scanner.NewScanner(verbose)
		sc.Scan(bh, output, workers)
	},
}

func init() {

	scanCmd.Flags().StringP("historyFile", "f", "", "path to history file")
	scanCmd.MarkFlagRequired("historyFile")
	scanCmd.Flags().StringP("output", "o", "report.md", "output file")
	scanCmd.Flags().BoolP("verbose", "v", false, "enable verbose checks. Might generate noisy and false positive findings")

	// searchCmd.Flags().StringP("check", "c", "", "run a specific check")

	rootCmd.AddCommand(scanCmd)

}

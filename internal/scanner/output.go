package scanner

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kmarkela/wiggumizer/internal/scanner/splugin"
)

func (s *Scanner) saveReport(fname string, findings map[string][]splugin.Finding, scope []string) error {

	content := "# Wiggumizer Report\n\n"
	content += "__Scope:__\n"

	for _, host := range scope {
		content += "- " + host + "\n"
	}
	content += "\n\n"
	content += "__List of Checks:__\n"

	for key, val := range s.checkers {
		content += "- __" + key + ":__ " + val.GetDescr() + "\n"
	}

	content += strings.Repeat("-", 30)
	content += "\n\n"

	for key, val := range findings {

		content += "## " + key + "\n"

		for i, f := range val {
			content += "###  #" + strconv.Itoa(i) + " - " + f.Description + "\n"
			content += "__Host: " + f.Host + "__ \n\n"
			content += "_Details:_\n\n```\n" + f.Evidens + "\n```\n"

			if f.Details != "" {
				content += "_Additional Details:_\n\n```\n" + f.Details + "\n```\n"
			}
		}

	}

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	file.WriteString(content)

	fmt.Printf("Scan report saved to: %s\n", fname)

	return nil

}

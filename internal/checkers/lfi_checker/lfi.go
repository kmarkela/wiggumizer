package lfi_checker

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
	"gopkg.in/yaml.v2"
)

var configFile string = "internal/checkers/lfi_checker/config.yaml"

type lfiChecker struct {
	name, descr string
}

// declare checker
func New() lfiChecker {
	return lfiChecker{
		name:  "lfi_checker",
		descr: "This module is searching for filenames in request parameters. Could be an indication of possible LFI",
	}
}

func (lc lfiChecker) GetName() string {
	return lc.name
}

func (lc lfiChecker) GetDescr() string {
	return lc.descr
}

func getExtList() ([]string, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return []string{}, err
	}

	var extensions []string
	if err := yaml.Unmarshal(file, &extensions); err != nil {
		return []string{}, err
	}

	return extensions, nil
}

func (lc lfiChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding
	var rePatern = ".*\\.("

	extL, err := getExtList()
	if err != nil {
		log.Printf("Cannot get list of Ext for LFI. Err: %s\n", err.Error())
		return f, false
	}

	for i, ext := range extL {
		//`.*\.(txt|php|exe)$`
		rePatern = rePatern + strings.ReplaceAll(ext, ".", "") // add exts w\o leading dot
		if i == len(extL)-1 {
			rePatern = rePatern + ").*"
			break
		}
		rePatern = rePatern + "|"

	}

	regex, err := regexp.Compile(rePatern)
	if err != nil {
		return f, false
	}

	var getParams string
	pParts := strings.SplitN(hi.Path, "?", 1)
	if len(pParts) == 2 {
		getParams = pParts[1]
	}

	if match := regex.MatchString(hi.Req.Body + getParams); !match {
		return f, false
	}

	f = splugin.Finding{
		Host:        hi.Host,
		Description: "filename found in Req params",
		Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),
	}

	if len(hi.Req.Body) != 0 {
		f.Details = hi.Req.Body
	}

	return f, true

}

package secretchecker

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kmarkela/wiggumizer/internal/historyparser"
	"github.com/kmarkela/wiggumizer/internal/scanner/splugin"
	"gopkg.in/yaml.v2"
)

var configFile string = "internal/checkers/secret_checker/config.yaml"

type RulesList struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Description string `yaml:"description"`
	Regex       string `yaml:"regex"`
	Verbose     bool
}

type secretChecker struct {
	name, descr string
	verbose     bool
}

// declare checker
func New() splugin.Checker {
	return &secretChecker{
		name:  "secret_checker",
		descr: "This module is searching for secrets (eg. API keys)",
	}
}

func (sc secretChecker) GetName() string {
	return sc.name
}

func (sc secretChecker) GetDescr() string {
	return sc.descr
}

func (sc secretChecker) GetVerbose() bool {
	return sc.verbose
}

func (sc *secretChecker) SetVerbose(v bool) error {
	sc.verbose = v
	return nil
}

func getRulesList() ([]Rule, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return []Rule{}, err
	}

	var rList RulesList
	if err := yaml.Unmarshal(file, &rList); err != nil {
		return []Rule{}, err
	}

	return rList.Rules, nil
}

func (sc secretChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	rl, err := getRulesList()
	if err != nil {
		log.Printf("Cannot get list of regex. Err: %s\n", err.Error())
		return f, false
	}
	for _, rule := range rl {

		// skip verbose rules if vervose is not set for the check
		if rule.Verbose && !sc.verbose {
			continue
		}

		regex, err := regexp.Compile(strings.TrimSuffix(rule.Regex, "\n"))
		if err != nil {
			log.Printf("Error compiling regex pattern: %s\n", err)
			continue
		}

		match := regex.FindString(hi.Req.Body + hi.Req.Headers + hi.Res.Body + hi.Res.Headers)

		if match == "" {
			continue
		}

		f.Host = hi.Host
		f.Description = "Secrets Found"
		f.Evidens = fmt.Sprintf("Path: %s", hi.Path)
		f.Details = f.Details + fmt.Sprintf("Description: %s\n Match: %s\n", rule.Description, match)

	}

	if f.Host != "" {
		return f, true
	}

	return f, false
}

package secretchecker

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

var configFile string = "internal/checkers/secret_checker/config.yaml"

type RulesList struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Description string `yaml:"description"`
	Regex       string `yaml:"regex"`
}

type secretChecker struct {
	name, descr string
}

// declare checker
func New() secretChecker {
	return secretChecker{
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

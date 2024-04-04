package subdchecker

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
	"gopkg.in/yaml.v2"
)

var configFile string = "internal/checkers/subd_checker/config.yaml"

type NotFoundServices struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Name            string `yaml:"name"`
	NotFoundMessage string `yaml:"notFoundMessage"`
}

type subdChecker struct {
	name, descr string
	verbose     bool
}

// declare checker
func New() splugin.Checker {
	return &subdChecker{
		name:  "subdomain_checker",
		descr: "This module is searching for 404 messages form hosting platformas",
	}
}

func (sc subdChecker) GetName() string {
	return sc.name
}

func (sc subdChecker) GetDescr() string {
	return sc.descr
}

func (sc subdChecker) GetVerbose() bool {
	return sc.verbose
}

func (sc *subdChecker) SetVerbose(v bool) error {
	sc.verbose = v
	return nil
}

func getNFList() ([]Service, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return []Service{}, err
	}

	var nfList NotFoundServices
	if err := yaml.Unmarshal(file, &nfList); err != nil {
		return []Service{}, err
	}
	return nfList.Services, nil
}

func (sc subdChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	lnf, err := getNFList()
	if err != nil {
		log.Printf("Cannot get list of not found messages. Err: %s\n", err.Error())
		return f, false
	}

	if hi.Status != 404 {
		return f, false
	}

	for _, nfm := range lnf {
		if !strings.Contains(hi.Res.Body, nfm.NotFoundMessage) {
			continue
		}

		f = splugin.Finding{
			Host:        hi.Host,
			Description: fmt.Sprintf("404 Message from %s found", nfm.Name),
			Evidens:     fmt.Sprintf("Path: %s", hi.Path),
		}
		return f, true

	}

	return f, false
}

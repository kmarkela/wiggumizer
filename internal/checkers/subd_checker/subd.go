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

type notFoundMessage struct {
	name, notFoundMessage string
}

type subdChecker struct {
	name, descr string
}

// declare checker
func New() subdChecker {
	return subdChecker{
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

func getNFList() ([]notFoundMessage, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return []notFoundMessage{}, err
	}

	var nfList []notFoundMessage
	if err := yaml.Unmarshal(file, &nfList); err != nil {
		return []notFoundMessage{}, err
	}

	return nfList, nil
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
		if !strings.Contains(hi.Res.Body, nfm.notFoundMessage) {
			continue
		}

		f = splugin.Finding{
			Host:        hi.Host,
			Description: fmt.Sprintf("404 Message from %s found", nfm.name),
			Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),
		}
		break

	}

	return f, true
}

package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/kmarkela/Wiggumizeng/pkg/collections"
)

func GetMUltipleChoices(label string, opts []string) collections.Set {
	var res = []string{}
	var hs = collections.Set{}

	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opts,
		PageSize: 15,
	}
	survey.AskOne(prompt, &res)

	for _, item := range res {
		hs.Add(item)
	}

	return hs
}

func GetString(msg string) string {
	var s string

	prompt := &survey.Input{
		Message: msg,
	}
	survey.AskOne(prompt, &s)

	return s
}

func GetSelect(msg string, opts []string, def string) string {
	var s string

	prompt := &survey.Select{
		Message: msg,
		Options: opts,
		Default: def,
	}
	survey.AskOne(prompt, &s)

	return s

}

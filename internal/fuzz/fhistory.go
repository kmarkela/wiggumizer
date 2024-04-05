package fuzz

import (
	"fmt"

	"github.com/kmarkela/Wiggumizeng/pkg/collections"
)

type fhistory struct {
	h map[string]collections.Set
}

func (fh *fhistory) add(s string) {
	if _, ok := fh.h[s]; !ok {
		fh.h[s] = collections.Set{}
	}
}

func (fh *fhistory) get(s string) (collections.Set, error) {
	if val, ok := fh.h[s]; ok {
		return val, nil
	}
	return collections.Set{}, fmt.Errorf("endpoint not found")
}

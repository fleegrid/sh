package sh

import (
	"strings"
	"testing"
)

func TestSh(t *testing.T) {

	tpl := `
	#!/bin/sh
	NAME={{.Name}}
	echo $NAME
	`

	params := Params{
		"Name": "DUST",
	}

	s, err := Run(tpl, params)

	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(s) != "DUST" {
		t.Errorf("%v != DUST", s)
	}
}

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

func TestExtractResult(t *testing.T) {
	tpl := `
	#!/bin/sh

	echo "----------"
	echo "hello"
	echo "----------"
	echo " world"
	`
	s, err := Run(tpl, nil)

	if err != nil {
		t.Error(err)
	}

	r := ExtractResult(s)

	if r != "world" {
		t.Error("not equal")
	}
}

func TestExtractResultMultiline(t *testing.T) {
	tpl := `
	#!/bin/sh

	echo "----------"
	echo "hello"
	echo "----------"
	echo " world"
	echo "yes "
	echo "  ho  "
	`
	s, err := Run(tpl, nil)

	if err != nil {
		t.Error(err)
	}

	r := ExtractResult(s)

	if r != "world\nyes\nho" {
		t.Error("not equal")
	}
}

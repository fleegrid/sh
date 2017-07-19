package sh

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
	"text/template"
)

// Params is an alias to map[string]interface{}, sh.Params is much shorter
type Params map[string]interface{}

// Run executes shell with template and params
func Run(tpl string, p interface{}) (out string, err error) {
	var t *template.Template
	var stdin io.WriteCloser
	var b []byte

	// prepare the command
	c := exec.Command("sh", "-s")

	// extract stdin
	if stdin, err = c.StdinPipe(); err != nil {
		return
	}

	// prepare the template
	if t, err = template.New("").Parse(tpl); err != nil {
		return
	}

	// connect template and Cmd.Stdin
	go func() {
		defer stdin.Close()
		t.Execute(stdin, p)
	}()

	// extract output
	if b, err = c.CombinedOutput(); err != nil {
		return
	}

	out = string(b)
	return
}

// ExtractResult extract lines below "------" (6 dash) as command result
func ExtractResult(out string) string {
	ret := ""
	br := bufio.NewReader(strings.NewReader(out))

	mark := false

	for {
		l, _, err := br.ReadLine()
		// error occurred, basically EOF
		if err != nil {
			break
		}
		s := strings.TrimSpace(string(l))
		// mark 6-dash
		if strings.Contains(s, "------") {
			if mark {
				// clear ret
				ret = ""
			} else {
				// mark
				mark = true
			}
		} else {
			if mark {
				// append line
				if len(ret) > 0 {
					ret += "\n" + s
				} else {
					ret = s
				}
			}
		}
	}
	return ret
}

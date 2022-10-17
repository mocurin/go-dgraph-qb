package gdqb

import (
	"fmt"
	"strings"

	"github.com/mocurin/go-dgraph-qb/build"
)

type Query interface {
	ComposeQuery(minified bool) ([]string, build.BuildSummary, error)
}

type Q struct {
	Name  string
	Args  []Argument
	Dirs  []Directive
	Preds []Predicate
}

func (q *Q) ComposeQuery(minified bool) (lines []string, summary build.BuildSummary, err error) {
	// Queries must always have a name
	if q.Name == "" {
		err = fmt.Errorf("query `Name` field is required")

		return
	}

	// Compose query arguments
	args, _, err := ComposeArgs(q.Args, minified)

	if err != nil {
		return
	}

	// Every function has at least one argument
	if len(args) == 0 {
		err = fmt.Errorf("queries always have at least one argument, got none")

		return
	}

	// Compose query directives
	dirs, _, err := ComposeDirs(q.Dirs, minified)

	if err != nil {
		return
	}

	// Compose query predicates
	preds, _, err := ComposePreds(q.Preds, minified)

	if err != nil {
		return
	}

	if len(preds) == 0 {
		err = fmt.Errorf("queries must have a selection of subfields, got none")
	}

	if minified {
		args := args[0]

		line := []string{q.Name, "(", args, ")"}

		if len(dirs) != 0 {
			dir := dirs[0]

			line = append(line, " ", dir)
		}

		preds := preds[0]

		line = append(line, " { ", preds, " }")

		lines = []string{strings.Join(line, "")}

		return
	}

	lines = make([]string, 0, len(args)+len(preds)+3)

	lines = append(lines, q.Name+"(\n")
	lines = append(lines, Identate(args)...)

	if len(dirs) != 0 {
		dirs := dirs[0]

		lines = append(lines, ") ", dirs, " {\n")
	} else {
		lines = append(lines, ") {")
	}

	lines = append(lines, Identate(preds)...)
	lines = append(lines, "}\n")

	return
}

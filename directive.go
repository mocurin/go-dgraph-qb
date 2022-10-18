package gdqb

import (
	"fmt"
	"strings"

	"github.com/mocurin/go-dgraph-qb/build"
	"github.com/mocurin/go-dgraph-qb/dqld"
)

type Directive interface {
	ComposeDirective(minified bool) ([]string, build.BuildSummary, error)
}

func ComposeDirs(dirs []Directive, minified bool) (lines []string, summary build.BuildSummary, err error) {
	composed := make([]string, 0, len(dirs))

	// Compose every directive
	for _, dir := range dirs {
		// TODO(@mocurin): Ignore build info for now
		dir, _, derr := dir.ComposeDirective(minified)

		// Break on the first encountered error
		if derr != nil {
			err = derr

			return
		}

		// Merge last & first lines in composed directive
		if !minified {
			if len(composed) != 0 {
				last := len(composed) - 1

				composed[last] = fmt.Sprintf("%s %s", composed[last], dir[0])
				dir = dir[1:]
			}
		}

		composed = append(composed, dir...)
	}

	if minified {
		line := strings.Join(composed, " ")

		lines = []string{line}

		return
	}

	lines = composed

	return
}

type D struct {
	Type dqld.DirectiveType
	Args []Argument
}

func (d *D) ComposeDirective(minified bool) (lines []string, summary build.BuildSummary, err error) {
	lines = []string{fmt.Sprintf("@%s", d.Type)}

	if len(d.Args) == 0 {
		return
	}

	args, _, err := ComposeArgs(d.Args, minified)

	if err != nil {
		return
	}

	if minified {
		lines = append(lines, "(", args[0], ")")

		line := strings.Join(lines, "")

		lines = []string{line}

		return
	}

	lines[0] += "(\n"
	lines = append(lines, Indentate(args)...)
	lines = append(lines, ")")

	return
}

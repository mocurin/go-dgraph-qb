package gdqb

import (
	"strings"

	"github.com/mocurin/go-dgraph-qb/build"
	"github.com/mocurin/go-dgraph-qb/dqlt"
)

type Argument interface {
	ComposeArgument(minified bool) ([]string, build.BuildSummary, error)
}

func ComposeArgs(args []Argument, minified bool) (lines []string, summary build.BuildSummary, err error) {
	composed := make([]string, 0, len(args))

	// Compose every argument
	for idx, arg := range args {
		if arg == nil {
			return
		}

		// TODO(@mocurin): Ignore build info for now
		arg, _, aerr := arg.ComposeArgument(minified)

		// Break on the first encountered error
		if aerr != nil {
			err = aerr

			return
		}

		// Add line break to the last one only
		// Argument composer takes care of all possibly nested data
		if !minified {
			last := len(arg) - 1

			// Do not add comma on last line to avoid syntax error
			if idx != len(args)-1 {
				arg[last] += ","
			}

			arg[last] += "\n"
		}

		composed = append(composed, arg...)
	}

	// On minified mode use a single line representation
	if minified && len(composed) != 0 {
		line := strings.Join(composed, ", ")

		lines = []string{line}

		return
	}

	lines = composed

	return
}

type A struct {
	Value interface{}
	Type  dqlt.DataType
}

type NA struct {
	A
	Name string
}

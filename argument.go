package gdqb

import (
	"fmt"
	"strings"
	"time"

	"github.com/mocurin/go-dgraph-qb/build"
	"github.com/mocurin/go-dgraph-qb/dqlt"
)

type Argument interface {
	ComposeArgument(minified bool) ([]string, build.BuildSummary, error)
}

type NamedArgument interface {
	Argument
	ComposeNamedArgument(minified bool) ([]string, build.BuildSummary, error)
}

func combineArgs(args [][]string, minified bool) []string {
	composed := make([]string, 0, len(args))

	// Compose every argument
	for idx, arg := range args {
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

		return []string{line}
	}

	return composed
}

func ComposeArgs(args []Argument, minified bool) (lines []string, summary build.BuildSummary, err error) {
	composed := make([][]string, 0, len(args))

	// Compose every argument
	for _, arg := range args {
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

		composed = append(composed, arg)
	}

	lines = combineArgs(composed, minified)

	return
}

func ComposeNamedArgs(args []NamedArgument, minified bool) (lines []string, summary build.BuildSummary, err error) {
	composed := make([][]string, 0, len(args))

	// Compose every argument
	for _, arg := range args {
		if arg == nil {
			return
		}

		// TODO(@mocurin): Ignore build info for now
		arg, _, aerr := arg.ComposeNamedArgument(minified)

		// Break on the first encountered error
		if aerr != nil {
			err = aerr

			return
		}

		composed = append(composed, arg)
	}

	lines = combineArgs(composed, minified)

	return
}

type A struct {
	Value interface{}
	Type  dqlt.DataType
}

func (a A) ComposeArgument(minified bool) (lines []string, summary build.BuildSummary, err error) {
	var line string

	switch a.Type {
	case dqlt.Geo:
		err = fmt.Errorf("geo value arguments are not implemented yet")

		return
	case dqlt.DateTime:
		date, ok := a.Value.(time.Time)

		if !ok {
			err = fmt.Errorf("value %s in an argument is not DateTime object", a.Value)

			return
		}

		a.Value = date.Format(time.RFC3339)

		fallthrough
	case dqlt.String, dqlt.Password:
		line = fmt.Sprintf(`"%s"`, a.Value)
	default:
		line = fmt.Sprint(a.Value)
	}

	lines = []string{line}

	return
}

type NA struct {
	A
	Name string
}

func (na NA) ComposeNamedArgument(minified bool) (lines []string, summary build.BuildSummary, err error) {
	if na.Name == "" {
		err = fmt.Errorf("named argument is expected to have a name, got none")

		return
	}

	lines, _, err = na.ComposeArgument(minified)

	if err != nil {
		return
	}

	lines[0] = fmt.Sprintf("%s: %s", na.Name, lines[0])

	return
}

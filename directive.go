package gdqb

import (
	"strings"

	"github.com/mocurin/go-dgraph-qb/build"
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

		composed = append(composed, dir...)
	}

	// Do not add empty string in case got no directives
	if len(composed) == 0 {
		return
	}

	// Directives can not be compose in non-minfied way
	line := strings.Join(composed, " ")

	lines = []string{line}

	return
}

type D struct{}

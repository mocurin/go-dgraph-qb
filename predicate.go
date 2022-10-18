package gdqb

import (
	"fmt"
	"strings"

	"github.com/mocurin/go-dgraph-qb/build"
)

type Predicate interface {
	ComposePredicate(minified bool) ([]string, build.BuildSummary, error)
}

func ComposePreds(preds []Predicate, minified bool) (lines []string, summary build.BuildSummary, err error) {
	composed := make([]string, 0, len(preds))
	for _, pred := range preds {
		if pred == nil {
			continue
		}

		pred, _, perr := pred.ComposePredicate(minified)

		if perr != nil {
			err = perr

			return
		}

		// Add line break to the last one only
		// Predicate composer takes care of all possibly nested predicates
		if !minified {
			pred[len(pred)-1] += "\n"
		}

		composed = append(composed, pred...)
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

type P struct {
	Name  string
	Alias string
}

func (p *P) ComposePredicate(minified bool) (lines []string, summary build.BuildSummary, err error) {
	if p.Name == "" {
		err = fmt.Errorf("predicate is expected to have name, got none")
	}

	if p.Alias != "" {
		line := fmt.Sprintf("%s: %s", p.Alias, p.Name)

		lines = []string{line}

		return
	}

	lines = []string{p.Name}

	return
}

type PN struct {
	P
	Dirs  []Directive
	Preds []Predicate
}

func (pn *PN) ComposePredicate(minified bool) (lines []string, summary build.BuildSummary, err error) {
	lines, _, err = pn.P.ComposePredicate(minified)

	if err != nil {
		return
	}

	dirs, _, err := ComposeDirs(pn.Dirs, minified)

	if err != nil {
		return
	}

	preds, _, err := ComposePreds(pn.Preds, minified)

	if err != nil {
		return
	}

	if len(dirs) != 0 {
		dirs := dirs[0]

		lines = append(lines, " ", dirs)
	}

	if minified {
		preds := preds[0]

		lines = append(lines, "{ ", preds, " }")
		lines = []string{strings.Join(lines, "")}

		return
	}

	lines = append(lines, " {\n")
	lines = []string{strings.Join(lines, "")}
	lines = append(lines, Indentate(preds)...)
	lines = append(lines, "}")

	return
}

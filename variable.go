package gdqb

import "github.com/mocurin/go-dgraph-qb/build"

type Variable interface {
	ComposeVariable(minified bool) ([]string, build.BuildSummary, error)
}

type V struct{}

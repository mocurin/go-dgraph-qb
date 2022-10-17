package gdqb

import "github.com/mocurin/go-dgraph-qb/build"

type Function interface {
	ComposeFunction(minified bool) ([]string, build.BuildSummary, error)
}

type F struct{}

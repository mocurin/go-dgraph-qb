package gdqb

import (
	"github.com/mocurin/go-dgraph-qb/build"
	"github.com/mocurin/go-dgraph-qb/dqlf"
)

type Function interface {
	ComposeFunction(minified bool) ([]string, build.BuildSummary, error)
}

type F struct {
	Args []Argument
	Type dqlf.FunctionType
}

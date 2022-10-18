package gdqb

import (
	"strings"
	"testing"

	"github.com/mocurin/go-dgraph-qb/dqld"
	"github.com/mocurin/go-dgraph-qb/dqlt"
)

func TestSimpleQuery(t *testing.T) {
	q := Q{
		Name: "q",
		Args: []NamedArgument{
			&NA{
				Name: "first",
				A: A{
					Type:  dqlt.Integer,
					Value: 10,
				},
			},
			&NA{
				Name: "offset",
				A: A{
					Type:  dqlt.Integer,
					Value: 10,
				},
			},
		},
		Dirs: []Directive{
			&D{
				Type: dqld.Normalize,
			},
			&D{
				Type: dqld.Cascade,
			},
			&D{
				Type: dqld.Filter,
				Args: []Argument{
					&A{
						Type:  dqlt.String,
						Value: "123123",
					},
				},
			},
		},
		Preds: []Predicate{
			&P{
				Name:  "Port.id",
				Alias: "id",
			},
			&P{
				Name: "Port.first_seen",
			},
			&P{
				Name: "Port.last_seen",
			},
			&PN{
				P: P{
					Name: "Port.HAS_PORT_SERVICE",
				},
				Preds: []Predicate{
					&P{
						Name: "PortService.first_seen",
					},
					&P{
						Name: "PortService.last_seen",
					},
				},
			},
		},
	}

	qmax, _, err := q.ComposeQuery(false)

	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("\n%s", strings.Join(qmax, ""))
	}

	qmin, _, err := q.ComposeQuery(true)

	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("\n%s", qmin[0])
	}
}

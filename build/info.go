package build

// BuildSummary represents an accumultaion of important data recieved during query building
type BuildSummary interface {
	// EmbeddedVars returns data on variables expected to be embedded inside query string
	EmbeddedVars()
	// QueryVars returns data on variables expected to be passed as query arguments
	QueryVars()
}

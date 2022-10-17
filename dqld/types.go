package dqld

type DirectiveType string

const (
	Filter    DirectiveType = "filter"
	Normalize DirectiveType = "normalize"
	Cascade   DirectiveType = "cascade"
)

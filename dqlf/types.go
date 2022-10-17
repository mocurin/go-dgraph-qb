package dqlf

type FunctionType string

const (
	Uid            FunctionType = "uid"
	UidIn          FunctionType = "uid_in"
	Equal          FunctionType = "eq"
	Inequal        FunctionType = "ie"
	Between        FunctionType = "between"
	LessOrEqual    FunctionType = "le"
	LessThan       FunctionType = "lt"
	GreaterOrEqual FunctionType = "ge"
	GreaterThan    FunctionType = "gt"
	AllOfText      FunctionType = "alloftext"
	Regexp         FunctionType = "regexp"
	AnyOfTerms     FunctionType = "allofterms"
	AllOfTerms     FunctionType = "anyofterms"
	Has            FunctionType = "has"
	GeoNear        FunctionType = "near"
	GeoWithin      FunctionType = "within"
	GeoContains    FunctionType = "contains"
	GeoIntersects  FunctionType = "intersects"
)

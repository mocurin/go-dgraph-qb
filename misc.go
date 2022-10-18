package gdqb

var Indent string = "  "

func Indentate(lines []string) []string {
	indentated := make([]string, 0, len(lines))
	for _, line := range lines {
		indentated = append(indentated, Indent+line)
	}

	return indentated
}

package gdqb

var Ident string = ""

func Identate(lines []string) []string {
	identated := make([]string, 0, len(lines))
	for _, line := range lines {
		identated = append(identated, Ident+line)
	}

	return identated
}

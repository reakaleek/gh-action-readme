package markdown

import "strings"

func toc(lines []string, indent int, startDepth int) string {
	var sb strings.Builder
	for _, line := range lines {
		if isHeader(line) {
			headerLevel := getHeaderLevel(line)
			if headerLevel >= startDepth {
				numSpaces := (headerLevel - startDepth) * indent
				ind := strings.Repeat(" ", numSpaces)
				heading := strings.TrimLeft(line, "# ")
				sb.WriteString(ind + "- " + heading + "\n")
			}
		}
	}
	return strings.TrimSpace(sb.String())
}

func isHeader(line string) bool {
	return strings.HasPrefix(line, "#")
}

func getHeaderLevel(line string) int {
	level := 0
	for _, r := range line {
		if r == '#' {
			level++
		} else {
			break
		}
	}
	return level
}

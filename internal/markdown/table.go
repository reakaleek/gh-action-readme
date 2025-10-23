package markdown

import (
	"strings"
)

func table(matrix [][]string) string {
	if len(matrix) == 0 {
		return ""
	}
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	for i := 0; i < len(duplicate); i++ {
		for j := 0; j < len(duplicate[i]); j++ {
			duplicate[i][j] = strings.ReplaceAll(duplicate[i][j], "\n", "<br>")
		}
	}
	colWidths := getMaxLengths(duplicate)
	var sb strings.Builder
	attachHeader(&sb, duplicate, colWidths)
	attachBody(&sb, duplicate, colWidths)
	return sb.String()
}

func attachHeader(sb *strings.Builder, matrix [][]string, colWidths []int) {
	for i, col := range matrix[0] {
		sb.WriteString("| " + col + strings.Repeat(" ", colWidths[i]-len(col)+1)) // +1 for the leading space
	}
	sb.WriteString("|\n")
	for i := range matrix[0] {
		sb.WriteString("|" + strings.Repeat("-", colWidths[i]+2)) // +2 for the spaces
	}
	sb.WriteString("|\n")
}

func attachBody(sb *strings.Builder, matrix [][]string, colWidths []int) {
	for _, row := range matrix[1:] {
		sb.WriteString("|")
		for i, col := range row {
			sb.WriteString(" " + col + strings.Repeat(" ", colWidths[i]-len(col)) + " |")
		}
		sb.WriteString("\n")
	}
}

func getMaxLengths(data [][]string) []int {
	maxLengths := make([]int, len(data[0]))
	for _, row := range data {
		for i, col := range row {
			if len(col) > maxLengths[i] {
				maxLengths[i] = len(col)
			}
		}
	}
	return maxLengths
}

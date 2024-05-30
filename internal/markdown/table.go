package markdown

import (
	"strings"
)

func table(matrix [][]string) string {
	if len(matrix) == 0 {
		return ""
	}
	colWidths := getMaxLengths(matrix)
	var sb strings.Builder
	attachHeader(&sb, matrix, colWidths)
	attachBody(&sb, matrix, colWidths)
	return sb.String()
}

func attachHeader(sb *strings.Builder, matrix [][]string, colWidths []int) {
	for i, col := range matrix[0] {
		sb.WriteString("| " + col + strings.Repeat(" ", colWidths[i]-len(col)+1)) // +1 for the leading space
	}
	sb.WriteString("|\n")
	for i, _ := range matrix[0] {
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

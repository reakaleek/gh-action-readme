package helpers

import (
	"github.com/fatih/color"
)

// PrintSummary prints a colored summary line
// Example: "Summary: 5 updated, 3 unchanged"
func PrintSummary(count1 int, label1 string, color1 color.Attribute, count2 int, label2 string, color2 color.Attribute) {
	cyan := color.New(color.FgCyan)
	num1 := color.New(color1).SprintFunc()
	num2 := color.New(color2).SprintFunc()
	cyan.Printf("\nSummary: %s %s, %s %s\n", num1(count1), label1, num2(count2), label2)
}

// PrintHeader prints a cyan header line
// Example: "Found 10 action file(s)"
func PrintHeader(format string, args ...interface{}) {
	color.Cyan(format, args...)
}

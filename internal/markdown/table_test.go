package markdown

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable(t *testing.T) {
	// arrange
	matrix := [][]string{
		{"Name", "Description", "Required", "Default"},
		{"`input1`", "input1 description.", "`true`", "``"},
		{"`input2`", "input2 description longer.", "`false`", "`default value`"},
	}

	// act
	md := table(matrix)

	// assert
	assert.Equal(
		t,
		"| Name     | Description                | Required | Default         |\n"+
			"|----------|----------------------------|----------|-----------------|\n"+
			"| `input1` | input1 description.        | `true`   | ``              |\n"+
			"| `input2` | input2 description longer. | `false`  | `default value` |\n",
		md,
	)
}

func TestTableCellsShorterThanHeadings(t *testing.T) {
	// arrange
	matrix := [][]string{
		{"Name", "Description", "Required", "Default"},
		{"`i1`", "d1", "`true`", "` `"},
		{"`i2`", "d2", "`false`", "` `"},
	}

	// act
	md := table(matrix)

	// assert
	assert.Equal(
		t,
		"| Name | Description | Required | Default |\n"+
			"|------|-------------|----------|---------|\n"+
			"| `i1` | d1          | `true`   | ` `     |\n"+
			"| `i2` | d2          | `false`  | ` `     |\n",
		md,
	)
}

func TestEmpty(t *testing.T) {
	// arrange
	var matrix [][]string

	// act
	md := table(matrix)

	// assert
	assert.Equal(t, "", md)

}

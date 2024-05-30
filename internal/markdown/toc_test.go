package markdown

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToc(t *testing.T) {
	// arrange
	lines := []string{
		"# Foo",
		"## Bar",
		"### Baz",
		"## Qux",
	}

	// act
	result := toc(lines, 2, 2)

	// assert
	expected := strings.Join([]string{
		"- Bar",
		"  - Baz",
		"- Qux",
	}, "\n")
	assert.Equal(t, expected, result)
}

func TestToc2(t *testing.T) {
	// arrange
	lines := []string{
		"# Foo",
		"## Bar",
		"### Baz",
		"## Qux",
	}

	// act
	result := toc(lines, 3, 1)

	// assert
	expected := strings.Join([]string{
		"- Foo",
		"   - Bar",
		"      - Baz",
		"   - Qux",
	}, "\n")
	assert.Equal(t, expected, result)
}

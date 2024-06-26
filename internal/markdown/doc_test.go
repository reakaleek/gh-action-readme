package markdown

import (
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestDoc(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{
			"Hello,",
			"!",
		},
	}

	// act
	doc.insertAfterPrefix("Hello", "World")

	// assert
	assert.Equal(t, "Hello,\nWorld\n!", doc.ToString())
}

func TestInsertSection(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{
			"<!-- title -->",
			"World",
		},
	}

	// act
	doc.insertSection("title", "# Hello")
	// assert
	assert.Equal(t, "<!-- title -->\n# Hello\n<!--/title-->\nWorld", doc.ToString())
}

func TestRemoveSection(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{
			"<!-- title -->",
			"# Hello",
			"<!-- /title -->",
			"World",
		},
	}

	// act
	doc.clearSection("title")

	// assert
	assert.Equal(t, "<!-- title -->\nWorld", doc.ToString())
}

func TestDiff(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{"Hello, World"},
	}
	otherDoc := Doc{
		lines: []string{"Hello, World!"},
	}

	// act
	diff := doc.Diff(&otherDoc)

	// assert
	assert.Equal(t, "Hello, World\x1b[32m!\x1b[0m", diff.PrettyDiff)
}

func TestDiffTrue(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{"Hello, World"},
	}
	otherDoc := Doc{
		lines: []string{"Hello, World!"},
	}

	// act
	diff := doc.Diff(&otherDoc)

	// assert
	assert.True(t, diff.HasDiff)
}

func TestDiffFalse(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{"Hello, World!"},
	}
	otherDoc := Doc{
		lines: []string{"Hello, World!"},
	}

	// act
	diff := doc.Diff(&otherDoc)

	// assert
	assert.False(t, diff.HasDiff)
}

func TestSingleLineSection(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{
			"<!-- name --><!-- /name -->",
		},
	}

	// act
	doc.updateName("Foo")

	// assert
	assert.Equal(t, "<!-- name -->Foo<!-- /name -->", doc.ToString())
}

func TestGetAttribute(t *testing.T) {
	// arrange
	line := "<!--usage action=\"elastic/oblt-actions/test\" version=\"v2\"-->"

	// act
	action, _ := getAttribute(line, "action")
	version, _ := getAttribute(line, "version")

	// assert
	assert.Equal(t, "elastic/oblt-actions/test", action)
	assert.Equal(t, "v2", version)
}

func TestStartCommentPattern(t *testing.T) {
	// arrange
	pattern := startCommentPattern("usage")
	re := regexp.MustCompile(pattern)
	line := "<!--usage-->"

	// act
	result := re.MatchString(line)

	// assert
	assert.True(t, result)
}

func TestStartCommentPattern2(t *testing.T) {
	// arrange
	pattern := startCommentPattern("usage")
	re := regexp.MustCompile(pattern)
	line := "<!--usage action=\"action\" version=\"v1\"-->"

	// act
	result := re.MatchString(line)

	// assert
	assert.True(t, result)
}

func TestUpdateUsage(t *testing.T) {
	// arrange
	t.Setenv("VERSION", "v2")
	doc := Doc{
		lines: []string{
			"<!-- usage action=\"elastic/oblt-actions/test\" version=\"env:VERSION\" -->",
			"```yaml",
			"    uses: elastic/oblt-actions/test@v1",
			"```",
			"<!--/usage-->",
		},
	}

	// act
	err := doc.UpdateUsage(nil)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "<!-- usage action=\"elastic/oblt-actions/test\" version=\"env:VERSION\" -->\n```yaml\n    uses: elastic/oblt-actions/test@v2\n```\n<!--/usage-->", doc.ToString())
}

func TestUpdate(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{
			"<!--name--><!--/name-->",
			"<!--description-->",
			"<!--inputs-->",
			"<!--outputs-->",
			"<!-- usage action=\"elastic/oblt-actions/test\" version=\"v1\" -->",
			"```yaml",
			"    uses: elastic/oblt-actions/test@main",
			"```",
			"<!--/usage-->",
		},
	}

	a := action.New(
		"Test",
		"Author",
		"Test description.",
		action.Inputs{
			"input1": {
				Description: "input1 description.",
				Required:    true,
			},
		},
		[]string{"input1"},
		action.Outputs{
			"output1": {
				Description: "output1 description.",
			},
		},
		[]string{"output1"},
	)

	// act
	err := doc.Update(a)

	// assert
	assert.NoError(t, err)

	expected := strings.Join([]string{
		"<!--name-->Test<!--/name-->",
		"<!--description-->",
		"Test description.",
		"<!--/description-->",
		"<!--inputs-->",
		"| Name     | Description         | Required | Default |",
		"|----------|---------------------|----------|---------|",
		"| `input1` | input1 description. | `true`   | ` `     |",
		"<!--/inputs-->",
		"<!--outputs-->",
		"| Name      | Description          |",
		"|-----------|----------------------|",
		"| `output1` | output1 description. |",
		"<!--/outputs-->",
		"<!-- usage action=\"elastic/oblt-actions/test\" version=\"v1\" -->",
		"```yaml",
		"    uses: elastic/oblt-actions/test@v1",
		"```",
		"<!--/usage-->",
	}, "\n")

	assert.Equal(t, expected, doc.ToString())

}

func TestCopy(t *testing.T) {
	// arrange
	doc := Doc{
		lines: []string{"Hello, World"},
	}

	// act
	duplicate := doc.Copy()

	// assert
	assert.Equal(t, doc.ToString(), duplicate.ToString())

}

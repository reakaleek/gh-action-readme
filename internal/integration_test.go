package internal_test

import (
	"fmt"
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/reakaleek/gh-action-readme/internal/markdown"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	var tests = []struct {
		actionPath   string
		readmePath   string
		expectedPath string
		versionEnv   string
	}{
		{
			"testdata/1-action.yml",
			"testdata/1-README-in.md",
			"testdata/1-README-out.md",
			"v2.0.0",
		},
		{
			"testdata/2-action.yml",
			"testdata/2-README-in.md",
			"testdata/2-README-out.md",
			"v2.0.0",
		},
		{
			"testdata/3-action.yml",
			"testdata/3-README-in.md",
			"testdata/3-README-out.md",
			"v2.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s-%s", tt.actionPath, tt.readmePath, tt.expectedPath), func(t *testing.T) {
			t.Setenv("VERSION", tt.versionEnv)
			parser := action.NewParser()
			a, err := parser.Parse(tt.actionPath)
			if err != nil {
				t.Fatal(err)
			}
			doc, err := markdown.NewDoc(tt.readmePath)
			if err != nil {
				t.Fatal(err)
			}
			err = doc.Update(&a)
			if err != nil {
				t.Fatal(err)
			}
			updatedDoc, err := markdown.NewDoc(tt.expectedPath)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, updatedDoc.ToString(), doc.ToString())
		})
	}
}

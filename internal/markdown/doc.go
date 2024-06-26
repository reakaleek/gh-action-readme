package markdown

import (
	"errors"
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
	"regexp"
	"strings"
)

const (
	nameSectionName            = "name"
	descriptionSectionName     = "description"
	inputsSectionName          = "inputs"
	outputsSectionName         = "outputs"
	tableOfContentsSectionName = "toc"
)

type Doc struct {
	name  string
	lines []string
}

func NewDoc(name string) (*Doc, error) {
	content, err := readFile(name)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		newDoc := &Doc{
			name: name,
			lines: []string{
				"# <!--name--><!--/name-->",
				"<!--description-->",
				"## Inputs",
				"<!--inputs-->",
				"## Outputs",
				"<!--outputs-->",
				"## Usage",
				"<!--usage action=\"your/action\" version=\"v1\"-->",
				"```yaml",
				"on: push",
				"steps:",
				"  - uses: your/action@v1",
				"```",
				"<!--/usage-->",
			},
		}
		err = newDoc.WriteToFile()
		if err != nil {
			return nil, err
		}
		return newDoc, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return &Doc{
		name:  name,
		lines: strings.Split(content, "\n"),
	}, nil
}

func (d *Doc) updateName(name string) {
	d.clearSection(nameSectionName)
	d.insertSection(nameSectionName, name)
}

func (d *Doc) updateDescription(description string) {
	d.clearSection(descriptionSectionName)
	d.insertSection(descriptionSectionName, description)
}

func (d *Doc) updateInputs(inputsMatrix [][]string) {
	d.clearSection(inputsSectionName)
	d.insertSection(inputsSectionName, table(inputsMatrix))
}

func (d *Doc) updateOutputs(outputsMatrix [][]string) {
	d.clearSection(outputsSectionName)
	d.insertSection(outputsSectionName, table(outputsMatrix))
}

//func (d *Doc) updateTOC() {
//	d.clearSection(tableOfContentsSectionName)
//	d.insertSection(tableOfContentsSectionName, TableOfContents(d.lines))
//}

func (d *Doc) Update(a *action.Action) error {
	d.updateName(a.Name)
	d.updateDescription(a.Description)
	d.updateInputs(a.GetInputsMatrix())
	d.updateOutputs(a.GetOutputsMatrix())
	return d.UpdateUsage(a)
}

func (d *Doc) Copy() Doc {
	lines := make([]string, len(d.lines))
	copy(lines, d.lines)
	return Doc{
		name:  d.name,
		lines: lines,
	}
}

func (d *Doc) ToString() string {
	return strings.Join(d.lines, "\n")
}

func (d *Doc) WriteToFile() error {
	return os.WriteFile(d.name, []byte(d.ToString()), 0755)
}

func readFile(name string) (string, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func (d *Doc) findIndex(pattern string) int {
	r := regexp.MustCompile(pattern)
	for i, line := range d.lines {
		if r.MatchString(line) {
			return i
		}
	}
	return -1
}

func (d *Doc) insertAfterPrefix(prefix string, lines ...string) {
	index := d.findIndex(prefix)
	d.insertAfterIndex(index, lines...)
}

func (d *Doc) insertAfterIndex(index int, lines ...string) {
	d.lines = append(d.lines[:index+1], append(lines, d.lines[index+1:]...)...)
}

func (d *Doc) removeLines(start int, end int) {
	d.lines = append(d.lines[:start], d.lines[end:]...)
}

func startCommentPattern(name string) string {
	return fmt.Sprintf("<!--\\s*%s(\\s+\\w+=\"\\S+\")*\\s*-->", name)
}

func endCommentPattern(name string) string {
	return fmt.Sprintf("<!--\\s*\\/\\s*%s\\s*-->", name)
}

func insertBetweenMatches(str string, pattern1 string, pattern2 string, insertion string) string {
	re1 := regexp.MustCompile(pattern1)
	loc1 := re1.FindStringIndex(str)
	if loc1 == nil {
		return str
	}
	re2 := regexp.MustCompile(pattern2)
	loc2 := re2.FindStringIndex(str)
	if loc2 == nil {
		return str
	}
	return str[:loc1[1]] + insertion + str[loc2[0]:]
}

func (d *Doc) insertSection(name string, content string) {
	startIndex := d.findIndex(startCommentPattern(name))
	endIndex := d.findIndex(endCommentPattern(name))

	if startIndex == -1 {
		return
	}

	if startIndex == endIndex {
		d.lines[startIndex] = insertBetweenMatches(
			d.lines[startIndex],
			startCommentPattern(name),
			endCommentPattern(name),
			strings.TrimSpace(content),
		)
		return
	}
	d.insertAfterIndex(
		startIndex,
		strings.TrimSpace(content),
		fmt.Sprintf("<!--/%s-->", name),
	)
}

func (d *Doc) clearSection(name string) {
	endIndex := d.findIndex(endCommentPattern(name))
	if endIndex == -1 {
		return
	}
	startIndex := d.findIndex(startCommentPattern(name))
	d.removeLines(startIndex+1, endIndex+1)
}

func (d *Doc) GetName() string {
	return d.name
}

func (d *Doc) Diff(doc *Doc) DiffResult {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(d.ToString(), doc.ToString(), false)
	prettyDiff := dmp.DiffPrettyText(diffs)
	return DiffResult{
		PrettyDiff: prettyDiff,
		HasDiff:    len(diffs) > 1,
	}
}

type DiffResult struct {
	PrettyDiff string
	HasDiff    bool
}

func (d *Doc) UpdateUsage(a *action.Action) error {
	usageIndex := d.findIndex(startCommentPattern("usage"))
	usageEndIndex := d.findIndex(endCommentPattern("usage"))
	if usageIndex == -1 {
		return nil
	}
	if usageEndIndex == -1 {
		return fmt.Errorf("missing end comment for usage section. add <!--/usage--> to the end of the usage section")
	}
	version, err := getAttribute(d.lines[usageIndex], "version")
	if err != nil {
		return err
	}
	version, err = parseEnvVariable(version)
	if err != nil {
		return err
	}
	actionGlob, err := getAttribute(d.lines[usageIndex], "action")
	if err != nil {
		return err
	}
	versionedActionRe := regexp.MustCompile(`uses:\s*(\S+)@\S+`)
	if usageEndIndex > 0 && usageEndIndex > usageIndex {
		for i := usageIndex; i < usageEndIndex; i += 1 {
			submatch := versionedActionRe.FindStringSubmatch(d.lines[i])
			if len(submatch) == 2 {
				actionName := submatch[1]
				globMatch, _ := doublestar.Match(actionGlob, actionName)
				if globMatch {
					pattern := strings.ReplaceAll(fmt.Sprintf("%s@\\S+", actionName), "/", "\\/")
					re := regexp.MustCompile(pattern)
					d.lines[i] = re.ReplaceAllString(d.lines[i], fmt.Sprintf("%s@%s", actionName, version))
				}
			}
		}
	}
	return nil
}

func parseEnvVariable(variable string) (string, error) {
	if strings.HasPrefix(variable, "env:") {
		envVarName := strings.TrimPrefix(variable, "env:")
		envVarValue := os.Getenv(envVarName)
		if envVarValue == "" {
			return "", fmt.Errorf("the environment variable %s is not set", envVarName)
		}
		return envVarValue, nil
	} else {
		return variable, nil
	}
}

func getAttribute(line string, attribute string) (string, error) {
	pattern := regexp.MustCompile(fmt.Sprintf("<!--.*%s=\"(\\S*)\".*-->", attribute))
	matches := pattern.FindStringSubmatch(line)
	if len(matches) >= 2 {
		return matches[1], nil
	}
	return "", fmt.Errorf("failed to get attribute %s", attribute)
}

func (d *Doc) Equals(doc Doc) bool {
	return d.ToString() == doc.ToString()
}

<!-- Generated with https://github.com/reakaleek/gh-action-readme -->
# <!--name-->Test Action 2<!--/ name-->
<!--description-->Test Action 2 description.<!--/description-->

## Inputs
<!--  inputs  -->
| Name     | Description             | Required | Default         |
|----------|-------------------------|----------|-----------------|
| `input1` | input1 description. new | `true`   | ` `             |
| `input2` | input2 description. new | `false`  | `default value` |
<!--/inputs-->

## Outputs
<!--outputs-->
| Name      | Description              |
|-----------|--------------------------|
| `output1` | output1 description. new |
<!--/outputs-->

## Usage
<!--usage action="org/repo" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: org/repo@v2.0.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

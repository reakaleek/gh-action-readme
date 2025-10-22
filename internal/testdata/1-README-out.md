<!-- Generated with https://github.com/reakaleek/gh-action-readme -->
# <!--name-->Test Action<!--/ name-->
<!--description-->Test Action description.<!--/description-->

## Inputs
<!--  inputs  -->
| Name     | Description         | Required | Default         |
|----------|---------------------|----------|-----------------|
| `input1` | input1 description. | `true`   | ` `             |
| `input2` | input2 description. | `false`  | `default value` |
<!--/inputs-->

## Outputs
<!--outputs-->
| Name      | Description          |
|-----------|----------------------|
| `output1` | output1 description. |
<!--/outputs-->

## Usage
<!--usage action="org/repo/path" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses:  org/repo/path@v2.0.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

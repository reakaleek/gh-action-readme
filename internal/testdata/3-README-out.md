<!-- Generated with https://github.com/reakaleek/gh-action-readme -->
# <!--name-->Test Action<!--/ name-->
<!--description-->Test Action description.<!--/description-->

## Inputs
<!--  inputs  -->
| Name | Description | Required | Default |
|------|-------------|----------|---------|
<!--/inputs-->

## Outputs
<!--outputs-->
| Name | Description |
|------|-------------|
<!--/outputs-->

## Usage
<!--usage action="org/repo/test" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: org/repo/test@v2.0.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

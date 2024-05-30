# <!--name--> <!--/ name-->
<!--description--><!--/description-->

## Inputs
<!--  inputs  -->

## Outputs
<!--outputs-->

<!--/outputs-->

## Usage
<!--usage action="org/repo/path" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses:  org/repo/path@v0.9.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

# <!--name--> <!--/ name-->
<!--description--><!--/description-->

## Inputs
<!--  inputs  -->

## Outputs
<!--outputs-->

<!--/outputs-->

## Usage
<!--usage action="org/repo/test" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: org/repo/test@v0.9.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

<!-- Generated with https://github.com/reakaleek/gh-action-readme -->
## Usage
<!--usage action="org/repo/**/*" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: org/repo/tool/action-a@v1.0.0
      - uses: org/repo/action-b@v1.0.0 # Some comment
      - uses: org/repo/action-c@v1.0.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

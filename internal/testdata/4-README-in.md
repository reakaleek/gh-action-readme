## Usage
<!--usage action="org/repo/**/*" version="env:VERSION"-->
```yaml
jobs:
  job:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: org/repo/tool/action-a@0.12.0
      - uses: org/repo/action-b@0.15.0 # Some comment
      - uses: org/repo/action-c@0.3.0
        with:
          input1: value1
          input2: value2
```
<!--/usage-->

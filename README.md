# Installation
go install
# Usage
copy to clipboard stuff like

```
--- FAIL: TestOne (11.45s)
--- FAIL: TestTwo (22.89s)
--- FAIL: TestThree (11.15s)
```
and run `gtc` then you will find text `go test -v -run 'TestOne|TestTwo|TestThree'` in your clipboard.



# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	inspectionbase/cmd/inspectiond	[no test files]
ok  	inspectionbase/internal/api	0.018s
ok  	inspectionbase/internal/audit	0.021s
ok  	inspectionbase/internal/config	0.001s
ok  	inspectionbase/internal/domain	0.001s
--- FAIL: TestInspectionNotFoundError (0.02s)
    inspection_test.go:26: want not found, got <nil>
FAIL
FAIL	inspectionbase/internal/inspection	0.064s
ok  	inspectionbase/internal/reporting	0.002s
ok  	inspectionbase/internal/store	0.034s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/inspectiond): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/inspectiond): exit `0`

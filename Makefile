VERSION := 0.21.2
GO := CGO_ENABLED=0 GO111MODULE=on GOFLAGS="-count=1" go

-include .makerc

.PHONY: release
release:
	@env GOOS=linux GOARCH=amd64  ${GO} build -trimpath  -o ./bin/hash-report github.com/johnnywale/hash-report


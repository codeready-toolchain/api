.PHONY: lint
lint:
	$(Q)go get github.com/golangci/golangci-lint/cmd/golangci-lint
	$(Q)${GOPATH}/bin/golangci-lint ${V_FLAG} run


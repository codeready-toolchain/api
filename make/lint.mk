ifndef LINT_MK
LINT_MK:=# Prevent repeated "-include".

include ./make/verbose.mk
include ./make/go.mk

.PHONY: lint
lint:
	$(Q)go get github.com/golangci/golangci-lint/cmd/golangci-lint
	$(Q)${GOPATH}/bin/golangci-lint ${V_FLAG} run

endif


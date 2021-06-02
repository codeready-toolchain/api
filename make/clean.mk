.PHONY: clean
## Clean
clean: remove-bin remove-config
	$(Q)go clean ${X_FLAG} ./...

.PHONY: remove-bin
remove-bin:
	$(Q)rm -rf ./bin

.PHONY: remove-config
remove-config:
	$(Q)rm -rf config/

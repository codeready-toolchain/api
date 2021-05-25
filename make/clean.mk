.PHONY: clean
## Clean
clean: remove-vendor remove-config
	$(Q)go clean ${X_FLAG} ./...

.PHONY: remove-vendor
remove-vendor:
	$(Q)-rm -rf ${V_FLAG} ./vendor

.PHONY: remove-config
remove-config:
	$(Q)rm -rf config/

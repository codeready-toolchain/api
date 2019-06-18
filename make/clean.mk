.PHONY: clean
## Clean
clean:
	$(Q)-rm -rf ${V_FLAG} ./vendor
	$(Q)go clean ${X_FLAG} ./...
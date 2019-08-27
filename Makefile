.DEFAULT_GOAL: run

GO_PACKAGES := $(shell go list ./...)

.PHONY: run
run:
	@for p in $(GO_PACKAGES); do\
		echo "go run $$p"; \
		go run $$p; \
	done


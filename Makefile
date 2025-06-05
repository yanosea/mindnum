# shows help message defaultly
.DEFAULT_GOAL := help

#
# update
#
.PHONY: update.credits update.mocks update.swagger

# update `./CREDITS`
update.credits:
	gocredits -skip-missing . > ./CREDITS

# update mocks
update.mocks:
	# ./app/domain
	mockgen -source=./app/domain/mindnum/mindnum_repository.go -destination=./app/domain/mindnum/mindnum_repository_mock.go -package=mindnum
	# ./pkg/proxy
	mockgen -source=./pkg/proxy/buffer.go -destination=./pkg/proxy/buffer_mock.go -package=proxy
	mockgen -source=./pkg/proxy/cobra.go -destination=./pkg/proxy/cobra_mock.go -package=proxy
	mockgen -source=./pkg/proxy/debug.go -destination=./pkg/proxy/debug_mock.go -package=proxy
	mockgen -source=./pkg/proxy/os.go -destination=./pkg/proxy/os_mock.go -package=proxy
	mockgen -source=./pkg/proxy/pflag.go -destination=./pkg/proxy/pflag_mock.go -package=proxy

#
# test
#
.PHONY: test

# execute tests
test:
	@set -e; \
	if [ -f "./test.run" ]; then \
		echo "test already running"; \
		exit 1; \
	fi; \
	touch test.run; \
	go test -v -p 1 ./... -cover -coverprofile=./cover.out; \
	grep -v -E "(_mock\.go|/mock/|/proxy/|/docs/docs\.go)" ./cover.out > ./cover.out.tmp && mv ./cover.out.tmp ./cover.out; \
	go tool cover -html=./cover.out -o ./docs/coverage.html; \
	rm ./cover.out; \
	if [ -f "./test.run" ]; then \
		rm ./test.run; \
	fi

# required phony targets for standards
all: help
clean:
	@rm -f ./cover.out ./co
	@rm -f ./test.run ./con

# help
.PHONY: help
help:
	@echo ""
	@echo "available targets:"
	@echo ""
	@echo "  [update]"
	@echo "    update.credits       - update ./CREDITS file"
	@echo "    update.mocks         - update all mocks"
	@echo ""
	@echo "  [test]"
	@echo "    test                 - execute all tests in local"
	@echo ""

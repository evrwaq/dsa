.PHONY: test

test-ds:
	go test -count=1 -v ./tests/data_structures/linear
test-tree-based-ds:
	go test -count=1 -v ./tests/data_structures/tree-based
test-ds-coverage:
	go test -count=1 -coverprofile=coverage.out ./...

.PHONY: test

test:
	go test -count=1 -v ./tests/...
test-linear-ds:
	go test -count=1 -v ./tests/data_structures/linear/...
test-tree-based-ds:
	go test -count=1 -v ./tests/data_structures/tree_based/...
test-ds-coverage:
	go test -count=1 -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

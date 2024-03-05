coverage:
	go test ./... -coverprofile=coverage.out -json > test-execution-report.out && ./tools/scripts/exclude-from-coverage.sh
lint: 
	golangci-lint run --color always > lint-report.out
coverage:
	go test ./... -coverprofile=coverage.out -json > coverage-json-report.out && ./tools/scripts/exclude-from-coverage.sh
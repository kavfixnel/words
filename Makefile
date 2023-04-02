.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	gofmt -d ./

.PHONY: coverage
coverage:
	go test ./... --cover

.PHONY: words.cov
words.cov:
	go test ./... --coverprofile=words.cov

.PHONY: coverage-report
coverage-report: words.cov
	go tool cover --html=words.cov

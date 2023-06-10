tests:
	go test -v ./... -coverprofile coverage.out

coverage-html: tests ;
	go tool cover -html=coverage.out

coverage-console: tests ;
	go tool cover -func=coverage.out

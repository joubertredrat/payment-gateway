tests:
	go clean -testcache ./...
	go test -v ./... -coverprofile coverage.out

coverage-html: tests ;
	go tool cover -html=coverage.out

coverage-console: tests ;
	go tool cover -func=coverage.out

generate-mocks:
	mockgen -package=pkg -source=internal/domain/repository.go > pkg/domain_repository_mock.go
	mockgen -package=pkg -source=internal/domain/ports.go > pkg/domain_ports_mock.go
	mockgen -package=pkg -source=internal/application/dispatcher.go > pkg/application_dispatcher_mock.go
	mockgen -package=pkg -source=internal/application/logger.go > pkg/application_logger_mock.go

generate-protos:
	protoc --go-grpc_out=. --go_out=. --proto_path=../protos ../protos/authorization.proto

vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest --test ./...

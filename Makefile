build: 
	@go build -o bin/enfermeria_go

run: build
	@./bin/enfermeria_go

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migration-up:
	@go run cmd/migrate/main.go up

migration-down:
	@go run cmd/migrate/main.go down
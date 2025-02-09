run:
	go run main.go

test:
	go test -v ./...

test-single:
	go test -run $(name) -v ./...

migration:
	migrate create -ext sql -dir migrate/migrations $(name)

migrate-up:
	go run migrate/main.go up

migrate-down:
	go run migrate/main.go down

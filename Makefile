.PHONY: run test

run:
	go run main.go

test:
	go test -v -cover ./...

entity:
	go run -mod=mod entgo.io/ent/cmd/ent init ${name}
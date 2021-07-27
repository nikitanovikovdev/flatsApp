LINTER_IMAGE=golangci-lint:v1.41.1

.PHONY: lint
lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/${LINTER_IMAGE} golangci-lint run -v --timeout=20m --fix


push-flats:
	docker build -t belka256/flatsapp:latest .
	docker push belka256/flatsapp:latest

run:
	go run cmd/users/main.go
	go run cmd/web/main.go
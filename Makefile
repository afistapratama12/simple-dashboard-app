
setup:
	cd server && go mod download
	cd client && yarn install

start:
	docker-compose up -d db
	docker-compose up -d minio
	cd server && go run migration/migration.go

stop:
	docker-compose down

run-server:
	cd server && go run main.go

run-client:
	cd client && yarn dev

test-cover:
	cd server && go test -coverprofile=coverage.out ./...
	cd server && go tool cover -html=coverage.out
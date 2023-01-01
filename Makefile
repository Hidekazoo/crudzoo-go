test:
	go test -v ./e2e
dev:
	DB_SETTING='host=localhost port=5432 user=test dbname=test password=password sslmode=disable' go run main.go
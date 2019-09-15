project_id := ${PROJECT_ID}
version := ${GAE_VERSION}

start:
	go run main.go

migration:
	go run migration/migration.go

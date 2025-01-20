.PHONY: migrate_up, test_migrate_up, run

run:
	go run ./cmd/sso --config=./config/local.yaml

migrate_up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

test_migrate_up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations


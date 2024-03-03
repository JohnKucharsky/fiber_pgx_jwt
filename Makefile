dev:
	fresh
dev-db:
	docker compose -f compose-dev.yaml up -d
prod:
	docker compose up --build
tidy:
	go mod tidy
migrate:
	cd db/migrations; goose postgres postgres://postgres:pass@localhost:5432/data up
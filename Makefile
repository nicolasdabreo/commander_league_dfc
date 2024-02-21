run:
	bun run ./assets/index.ts --watch & \
	bunx tailwindcss -c './assets/tailwind.config.js' -i './assets/css/input.css' -o './assets/css/output.css' --watch & \
	air -c ./.air.toml

build:
	@go build -o _build/main cmd/app/main.go

migrate: 
	@go run ./cmd/migrate/main.go up

rollback:
	@go run ./cmd/migrate/main.go down

drop:
	@go run ./cmd/drop/main.go

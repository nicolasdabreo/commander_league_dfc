run:
	npx tailwind \
		-i './assets/css/input.css' \
		-o './assets/css/output.css' \
		--watch & \
	air -c ./.air.toml

build:
	@go build -o _build/main cmd/app/main.go

up: 
	@go run cmd/migrate/main.go up

down:
	@go run cmd/migrate/main.go down

seed: 
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

drop:
	@go run cmd/drop/main.go

reset:
	@go run cmd/seed/main.go


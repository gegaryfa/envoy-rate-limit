up:
	docker-compose up --force-recreate --build -d

down:
	docker-compose down

load-test:
	cd vegeta && go run main.go

.PHONY: up down load-test
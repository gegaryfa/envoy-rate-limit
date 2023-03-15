up:
	docker-compose up --force-recreate --build -d

down:
	docker-compose down

.PHONY: up down
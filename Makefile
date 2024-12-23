up: down
	docker compose up --build

down:
	docker compose down -v

clean: down

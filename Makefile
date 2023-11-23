.PHONY: init
init:
	docker compose run app go run ./
.PHONY: down
down:
	docker compose down

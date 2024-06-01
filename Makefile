.PHONY: setup up d b ps firebase build app

setup:
	@make up
	@make ps
d:
	docker compose down
up:
	docker compose up -d
ps:
	docker compose ps
firebase:
	docker compose exec firebase bash
app:
	docker compose exec app bash
build:
	docker compose build
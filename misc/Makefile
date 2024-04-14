.PHONY: start stop

start:
	docker compose -f compose.yaml up --build --force-recreate -d
	docker cp demo:/work/cli.json ../kubeshell/data/
stop:
	docker compose down
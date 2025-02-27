.PHONY: build run stop clean logs shell

SERVICE=app

build:
	docker-compose -f docker/docker-compose.yml build --no-cache

run:
	docker-compose -f docker/docker-compose.yml up --build --force-recreate

stop:
	docker-compose -f docker/docker-compose.yml down

clean:
	docker-compose -f docker/docker-compose.yml down --rmi all --volumes --remove-orphans

logs:
	docker-compose -f docker/docker-compose.yml logs -f $(SERVICE)

shell:
	docker-compose -f docker/docker-compose.yml run --rm $(SERVICE) sh

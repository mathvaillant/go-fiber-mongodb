.PHONY: start
start:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose rm -v --force --stop
	docker rmi go-fiber-mongodb

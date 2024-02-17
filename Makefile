.PHONY: up down clean all

up:
	docker-compose up -d

down:
	docker-compose down

clean: down
	docker rmi go-yinhuadmcc-server:latest

all: clean up
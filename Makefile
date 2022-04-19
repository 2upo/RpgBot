dev:
	docker-compose -f docker-compose.yml up --build web
test:
	docker-compose -f docker-compose.yml up --build test
